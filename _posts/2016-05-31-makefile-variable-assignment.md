---
layout: post
title:  "Makefile 中四种变量赋值的不同"
date:   2016-05-31 15:18:17 +0800
categories: make
keywords: make variable-assignment
---

[http://stackoverflow.com/questions/448910/makefile-variable-assignment](http://stackoverflow.com/questions/448910/makefile-variable-assignment)

## Lazy Set

    VARIABLE = value

当变量被用到的时候, 会递归的展开

## Imeediate Set

    VARIABLE := value

变量在申明的时候马上执行

## Set If Absent

    VARIABLE ?= value

变量不存在的时候才赋值(也是lazy set)

## Append

    VARIABLE += value

把value添加到已存在的变量后(如果变量不存在, 把value赋值给这个变量)


