'''
    modified from https://github.com/tuangauss/Various-projects/blob/master/Python/fireworks.py by Tuan Nguyen Doan
    changes were made to make firework more beautiful by changing particle trajectories
'''
import tkinter as tk
from PIL import Image, ImageTk
from time import time, sleep
from random import choice, uniform, randint
import numpy as np


GRAVITY = 100
 
colors = ['red', 'blue', 'yellow', 'white', 'green', 'orange', 'purple', 'seagreen', 'indigo', 'cornflowerblue']



class Particle:

    def __init__(self, cv, idx, total, explosion_speed, x=0., y=0., size=2., color='red', lifespan=2, delta=0, k=1,
                 **kwargs):
        self.id = idx
        self.x = x
        self.y = y
        self.initial_speed = explosion_speed
        self.total = total
        self.age = 0
        self.color = color
        self.delta = delta
        self.k = k
        self.cv = cv
        self.cid = self.cv.create_oval(
            x - size, y - size, x + size,
            y + size, fill=self.color)
        self.lifespan = lifespan

        self.flag = False
        self.beta_previouse = 0
        self.x_cordi_previouse = 0
        self.y_cordi_previouse = 0
        self.theta = 0
        self.beta_previouse = 0

    def update(self, dt):
        self.age += dt
        if self.alive() and self.expand():
            delta = self.delta * dt
            self.theta = self.theta + delta
            # used equation: hypotenuse = e^theta to calculate length from given angle(theta)
            hypotenuse = np.exp(self.k * self.theta)
            x_cordi = hypotenuse * np.cos(self.theta)
            y_cordi = hypotenuse * np.sin(self.theta)
            
            # initial velocity aiming at some certain direction according to idx
            alpha = np.radians(self.id * 360 / self.total)
            beta = np.arctan(y_cordi/x_cordi) # angle
            r = np.power(x_cordi ** 2 + y_cordi ** 2, 0.5)
            x_rotated = np.cos(alpha + beta) * r
            y_rotated = np.sin(alpha + beta) * r

            # ---------------------------- rotation to all directions --------------------------------
            # since arctan has domain between (-pi/2,pi/2), we need to consider situation out of range

            # when crossing pi, inverse symbol of x and y
            if self.beta_previouse < 0 and beta > 0:
                self.flag = not self.flag

            if self.flag:
                x_rotated = - x_rotated
                y_rotated = - y_rotated

            # if beta is in first second and fourth quadrants, inverse symbol of x and y
            if beta < 0:
                x_rotated = - x_rotated
                y_rotated = - y_rotated
            # ---------------------------- finished rotation part -------------------------------------

            self.beta_previouse = beta           
            move_x = x_rotated - self.x_cordi_previouse
            move_y = y_rotated - self.y_cordi_previouse

            self.x_cordi_previouse = x_rotated
            self.y_cordi_previouse = y_rotated

            self.vx = move_x / dt
            self.vy = move_y / dt
            self.cv.move(self.cid, move_x, move_y)

        elif self.alive():
            self.cv.move(self.cid, self.vx * dt, self.vy * dt + 1/2 * GRAVITY * dt)
            self.vy += GRAVITY * dt

        elif self.cid is not None:
            cv.delete(self.cid)
            self.cid = None

    def expand (self):
        return self.age <= 2.25

    def alive(self):
        return self.age <= self.lifespan



def simulate(cv):
    t = time()
    explode_points = []
    wait_time = randint(10, 100) # how long to wait before next simulation in milliseconds
    numb_explode = randint(6, 10)
    for point in range(numb_explode):
        objects = []
        x_cordi = randint(50, 550)
        y_cordi = randint(50, 150)
        size = uniform(0.5, 3)
        color = choice(colors)
        explosion_speed = uniform(0.2, 1)
        total_particles = randint(10, 50)
        for i in range(1, total_particles):
            # it's dirty but it works
            r = Particle(cv, idx=i, total=total_particles, explosion_speed=explosion_speed, x=x_cordi, y=y_cordi,
                         color=color, size=size, lifespan=uniform(2, 3.75), delta=1.8)
            r_ = Particle(cv, idx=i, total=total_particles, explosion_speed=explosion_speed, x=x_cordi, y=y_cordi,
                         color=color, size=size, lifespan=uniform(2, 3.75), delta=1.8, k=0.9)
            r__ = Particle(cv, idx=i, total=total_particles, explosion_speed=explosion_speed, x=x_cordi, y=y_cordi,
                         color=color, size=size, lifespan=uniform(2, 3.75), delta=1.8, k=0.788877)

            objects.append(r)
            objects.append(r_)
            objects.append(r__)
        explode_points.append(objects)
    total_time = .0
    while total_time < 4: # should be greater thatn lifespan of particles; otherwise some won't disappear
        sleep(0.01) # update every 0.01 second
        tnew = time()
        t, dt = tnew, tnew - t
        for point in explode_points:
            for item in point:
                item.update(dt)
        cv.update()
        total_time += dt
    root.after(wait_time, simulate, cv)


def close(*ignore):
    global root
    root.quit()


if __name__ == '__main__':
    root = tk.Tk()
    cv = tk.Canvas(root, height=400, width=600)
    image = Image.open("./night-sky.png")
    photo = ImageTk.PhotoImage(image)
    cv.create_image(0, 0, image=photo, anchor='nw')
    cv.pack()
    quitButton = tk.Button(root, text='Quit', command=cv.quit)
    quitButton.pack()

    root.protocol("WM_DELETE_WINDOW", close)
    root.after(100, simulate, cv)
    root.mainloop()
