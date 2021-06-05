from flask import Flask

PORT = 1802

app = Flask(__name__)

@app.route("/")
def index():
    return "Hello world"

if __name__ == "__main__":
    print("Python Flask server serving at port:", PORT)
    app.run(port=PORT)
