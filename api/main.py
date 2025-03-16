from fastapi import FastAPI, File, UploadFile
import uvicorn

app = FastAPI()

@app.get("/ping")
async def ping():
    return "Ping Successful"

@app.post("/predict")
async def predict(
file: UploadFile = File(...)
):
    bytes = await file.read()
    return


if __name__ == "__main__":
    uvicorn.run(app, host="localhost", port =3030)
