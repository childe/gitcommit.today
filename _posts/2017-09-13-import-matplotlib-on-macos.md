---

title: macos上面import matplotlib出错
date: 2017-09-13 14:42:12 +0800
layout: post
categories: python

---

报错:

    RuntimeError: Python is not installed as a framework.
    The Mac OS X backend will not be able to function correctly if Python is not installed as a framework.
    See the Python documentation for more information on installing Python as a framework on Mac OS X.
    Please either reinstall Python as a framework, or try one of the other backends.
    If you are using (Ana)Conda please install python.app and replace the use of ‘python’ with ‘pythonw’.
    See ‘Working with Matplotlib on OSX’ in the Matplotlib FAQ for more information.

解决:

    I assume you have installed the pip matplotlib,  
    there is a directory in you root called ~/.matplotlib.
    Create a file ~/.matplotlib/matplotlibrc there and add the following code: backend: TkAgg
