from flask import (
    Blueprint,
    render_template,
    request,
    redirect,
    url_for,
    session,
    jsonify,
)
from flask_login import login_user, logout_user, login_required
import requests
from src.models.user import User
from src.handlers.load_user import invalidate_user_cache
import os

auth_bp = Blueprint("auth", __name__, template_folder="templates")

BACKEND_URL = os.environ.get("BACKEND_URL", None)


@auth_bp.route("/login", methods=["GET", "POST"])
def login():
    if request.method == "GET":
        return render_template("login.html")
    elif request.method == "POST":
        user = User(**request.form.to_dict())
        login_response = requests.post(
            BACKEND_URL + "/login",
            json=user.to_dict(),
        )
        if login_response.status_code == 200:
            logged_in_user = User(**login_response.json()["data"])
            login_user(logged_in_user)
            response = {
                "status": "success",
                "message": login_response.json()["message"],
                "redirect_url": "/home",
            }
            return jsonify(response)

        else:
            response = {
                "status": "error",
                "message": login_response.json()["message"],
                "error": login_response.json()["error"],
            }
            return jsonify(response)


@auth_bp.route("/logout")
@login_required
def logout():
    logout_user()
    invalidate_user_cache()
    return redirect(url_for("home.index"))


@auth_bp.route("/register", methods=["GET", "POST"])
def register():
    if request.method == "GET":
        return render_template("register.html")
    elif request.method == "POST":
        user = User(**request.form.to_dict())
        register_response = requests.post(
            BACKEND_URL + "/create/user",
            json=user.to_dict(),
        )
        if register_response.status_code == 201:
            response = {
                "status": "success",
                "message": register_response.json()["message"],
                "redirect_url": "/login",
            }
            return jsonify(response)
        else:
            response = {
                "status": "error",
                "message": register_response.json()["message"],
                "error": register_response.json()["error"],
            }
            return jsonify(response)
