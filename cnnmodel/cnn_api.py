from urllib import response
import flask
from flask import request, jsonify
from numpy import append
import requests
import json
from requests.adapters import HTTPAdapter
from requests.packages.urllib3.util.retry import Retry
from sqlalchemy import true
import load_train1 as train
app = flask.Flask(__name__)
app.config["DEBUG"] = True


model = ""
classes = ["Adialer.C", "Agent.FYI", "Allaple.A", "Allaple.L",
           "Alueron.gen!J", "Autorun.K", "C2LOP.P", "C2LOP.gen!g",
           "Dialplatform.B", "Dontovo.A", "Fakerean", "Instantaccess", "Lolyda.AA1",
           "Lolyda.AA2", "Lolyda.AA3", "Lolyda.AT", "Malex.gen!J", "Obfuscator.AD", "Rbot!gen",
           "Skintrim.N", "Swizzor.gen!E", "Swizzor.gen!I", "VB.AT", "Wintrim.BX", "Yuner.A"]


@app.route('/container/cnnscan', methods=['OPTIONS'])
def cssScanOptions():
    response = flask.Response()
    response.headers["Access-Control-Allow-Origin"] = "*"
    response.headers["Access-Control-Allow-Credentials"] = True
    response.headers["Access-Control-Allow-Headers"] = "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Accept-Language,  Authorization, accept, origin, Cache-Control, X-Requested-With"
    response.headers["Access-Control-Allow-Methods"] = "POST, OPTIONS, GET, PUT , DELETE , PATCH"
    return response


@app.route('/container/cnnscan', methods=['GET'])
def cnnScanGet():
    if request.args.get("container") is None or request.args.get("path") is None:
        return jsonify({"success": "false", "message": "wrong contaienr"})

    '''resp = flask.Response()
    resp.headers["Access-Control-Allow-Origin"] = "*"
    resp.headers["Access-Control-Allow-Credentials"] = True
    resp.headers["Access-Control-Allow-Headers"] = "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Accept-Language,  Authorization, accept, origin, Cache-Control, X-Requested-With"
    resp.headers["Access-Control-Allow-Methods"] = "POST, OPTIONS, GET, PUT , DELETE , PATCH"'''
    response = ListContainerSpeceificFiles(
        request.args.get("container"), request.args.get("path"))
    response.headers.set("Access-Control-Allow-Origin", "*")


    response.headers.set("Access-Control-Allow-Credentials", "true")
    response.headers.set("Access-Control-Allow-Headers",
                      "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Accept-Language,  Authorization, accept, origin, Cache-Control, X-Requested-With")
    response.headers.set("Access-Control-Allow-Methods",
                      "POST, OPTIONS, GET, PUT , DELETE , PATCH")
    return response


def ListContainerSpeceificFiles(container, path):
    intializeExec = {"AttachStdin": False, "AttachStdout": True, "AttachStderr": True, "DetachKeys": "ctrl-p,ctrl-q", "Tty": False, "Cmd":
                     ["ls", "-a", path], "Env": ["FOO=bar", "BAZ=quux"]}

    startExec = {
        "Detach": False,
        "Tty": False
    }
    request = RunCommandsOnDockerContainer(
        container=container, intializeExec=intializeExec, startExec=startExec)

    request = GetFileContent(request.content.decode(
        "utf-8").split("\n"), container, path)
    print(request)
    return request


def RunCommandsOnDockerContainer(container, intializeExec, startExec):
    headers = {"Content-Type": "application/json"}

    request = requests.post(
        "http://172.17.0.1:5555/containers/"+container+"/exec", data=json.dumps(intializeExec), headers=headers, timeout=(2, 5))

    if request.status_code != 201:
        return request

    request = requests.post(
        "http://172.17.0.1:5555/exec/"+request.json()["Id"]+"/start", data=json.dumps(startExec), headers=headers, timeout=(2, 5))

    return request


def GetFileContent(files, container, path):
    if path[len(path)-1] != '/':
        path += "/"
    filesContent = []
    predictions = {"results": [{"passed": True,
                               "title": "CNN Scan", "details": []}], "compliance": 100}
    for i in range(2, len(files)-1):
        print(files[i])
        intializeExec = {"AttachStdin": False, "AttachStdout": True, "AttachStderr": True, "DetachKeys": "ctrl-p,ctrl-q", "Tty": False, "Cmd":
                         ["cat", path+files[i]], "Env": ["FOO=bar", "BAZ=quux"]}
        startExec = {
            "Detach": False,
            "Tty": False
        }
        request = RunCommandsOnDockerContainer(
            container, intializeExec, startExec)

        if request.status_code == 200:
            test = train.Predict(request.content)
            predictions["results"][0]["details"].append({"file": path+files[i],
                                           "description": classes[test[0]], "impact": ""})
    if len(predictions["results"][0]["details"]) !=0:
        predictions["results"][0]["passed"]=False
        predictions["compliance"] = 80
        request=StoreCnn(predictions)
        if request.status_code != 200:
            return request.body

    predictions["results"][0]["details"] = json.dumps(predictions["results"][0]["details"])
    return jsonify(predictions)

def StoreCnn(data):
    headers = {"Content-Type": "application/json"}

    reauest=requests.post("http://172.17.0.3:8080/container/storecnn",headers=headers,data=data)

    return request

if __name__ == '__main__':
    train.LoadModel()
    app.run(host="0.0.0.0", port=5001)
