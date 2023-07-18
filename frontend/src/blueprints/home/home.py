from flask import Blueprint, render_template
from flask_login import login_required

home_bp = Blueprint("home", __name__, template_folder="templates")


@home_bp.route("/")
def index():
    return render_template("index.html")


@home_bp.route("/home")
@login_required
def home():
    return render_template("home.html")
