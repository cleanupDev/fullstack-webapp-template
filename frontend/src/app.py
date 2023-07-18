from flask import Flask, render_template
from flask_login import LoginManager
import requests
from src.handlers.load_user import load_user
from src.blueprints.auth.auth import auth_bp
from src.blueprints.home.home import home_bp

app = Flask(__name__)
app.config["SECRET_KEY"] = "secret-key"
login_manager = LoginManager()
login_manager.init_app(app)
login_manager.user_loader(load_user)

app.register_blueprint(auth_bp)
app.register_blueprint(home_bp)


# Test route to check if frontend is able to communicate with backend
@app.route("/backend")
def api():
    return requests.get("http://backend:5001").content


@app.errorhandler(404)
def page_not_found(e):
    return render_template("404.html"), 404
