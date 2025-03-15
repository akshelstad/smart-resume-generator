from flask import Flask, request, jsonify
from openai_handler import generate_resume
import os
from dotenv import load_dotenv

load_dotenv()

app = Flask(__name__)

@app.route("/generate-resume", methods=["POST"])
def resume():
    data = request.get_json()
    name = data.get("name", "")
    experience = data.get("experience", "")
    resume_text = generate_resume(name, experience)
    return jsonify({"resume": resume_text})

if __name__ == "__main__":
    app.run(port=5001, debug=True)