from matplotlib import image
import tensorflow as tf
import numpy as np
from math import sqrt, ceil
import cv2

model = ""


def Predict(data):
    ConvertFileToImage(data)
    return RunModel()


def ConvertFileToImage(data):

    # Input file name (random file I found in my folder).

    #input_file_name = '/home/chall/College/four/gruation_project/archive/29c7e87350cb03428fc108b03856095b'

    # Read the whole file to data
    #with open(input_file_name, 'rb') as binary_file:
    data = data

    # Data length in bytes
    data_len = len(data)

    # d is a verctor of data_len bytes
    d = np.frombuffer(data, dtype=np.uint8)

    # Assume image shape should be close to square
    sqrt_len = int(ceil(sqrt(data_len)))  # Compute square toot and round up

    # Requiered length in bytes.
    new_len = sqrt_len*sqrt_len

    # Number of bytes to pad (need to add zeros to the end of d)
    pad_len = new_len - data_len

    # Pad d with zeros at the end.
    # padded_d = np.pad(d, (0, pad_len))
    padded_d = np.hstack((d, np.zeros(pad_len, np.uint8)))

    # Reshape 1D array into 2D array with sqrt_len pad_len x sqrt_len (im is going to be a Grayscale image).
    im = np.reshape(padded_d, (sqrt_len, sqrt_len))

    # Save image
    cv2.imwrite('im.png', im)
    RunModel()


def LoadModel():
    global model
    model = tf.keras.models.load_model('train1_saved')
    return model


def RunModel():
    global model
    img = "im.png"
    image_data = tf.keras.preprocessing.image.load_img(
        img, target_size=(64, 64))
    input_arr = tf.keras.preprocessing.image.img_to_array(image_data)

    predict = model.predict(np.array([input_arr])/255.)

    y_pred = np.argmax(predict, axis=1)

    return y_pred
