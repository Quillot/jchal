from PIL import Image
import os
for f in os.listdir("images"):
	im = Image.open("images/" + f)
	im.save(f + ".png")