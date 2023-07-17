from flask import (
    Flask,
    render_template,
    request,
    redirect,
    url_for,
    session,
    jsonify,
)
from flask_login import (
    LoginManager,
    login_user,
    logout_user,
    current_user,
    login_required,
)
import requests
import os

app = Flask(__name__)
app.config["SECRET_KEY"] = "secret-key"
login_manager = LoginManager()
login_manager.init_app(app)

backend_url = os.environ.get("BACKEND_URL", None)


class User:
    def __init__(self, id):
        self.id = id

    def get_id(self):
        return self.id

    def is_authenticated(self):
        return True

    def is_active(self):
        return True

    def is_anonymous(self):
        return False


@login_manager.user_loader
def load_user(user_id):
    user = User(user_id)
    return user


@app.route("/")
def index():
    return render_template("index.html")


@app.route("/login", methods=["GET", "POST"])
def login():
    if request.method == "GET":
        return render_template("login.html")
    elif request.method == "POST":
        user_data = request.form

        login_response = requests.post(
            backend_url + "/login",
            json={"username": user_data["username"], "password": user_data["password"]},
        )

        if login_response.status_code == 200:
            user = User(login_response.json()["id"])
            login_user(user)
            response = {
                "status": "success",
                "message": login_response.json()["message"],
                "user": login_response.json(),
            }["user"]
            return jsonify(response)
        else:
            # TODO: Get correct error message from backend
            response = {"status": "error", "message": login_response.json()["error"]}


@app.route("/logout")
@login_required
def logout():
    logout_user()
    return redirect(url_for("index"))


@app.route("/register", methods=["GET", "POST"])
def register():
    if request.method == "GET":
        return render_template("register.html")
    elif request.method == "POST":
        user_data = request.form

        register_response = requests.post(
            backend_url + "/create/user",
            json={
                "username": user_data["username"],
                "email": user_data["email"],
                "password": user_data["password"],
                "first_name": user_data["first_name"],
                "last_name": user_data["last_name"],
            },
        )

        if register_response.status_code == 200:
            user = User(register_response.json()["id"])
            login_user(user)
            response = {
                "status": "success",
                "message": register_response.json()["message"],
                "user": register_response.json()["user"],
            }
            return jsonify(response)
        else:
            response = {
                "status": "error",
                "message": register_response.json()["message"],
            }
            return jsonify(response)


@app.route("/protected")
@login_required
def protected():
    return "Works!"


@app.route("/backend")
def api():
    return requests.get("http://backend:5001").content


@app.errorhandler(404)
def page_not_found(e):
    return render_template("404.html"), 404
