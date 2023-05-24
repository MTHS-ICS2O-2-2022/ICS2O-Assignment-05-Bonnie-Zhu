// Copyright (c) Bonnie Zhu
//
// Created by: Bonnie Zhu
// Created on: May 2023
// This file contains the JS functions for index.html
"use strict"

function generatePyramid() {
  var inputNumber = parseInt(document.getElementById("inputNumber").value)
  var resultElement = document.getElementById("result")

  if (isNaN(inputNumber) || inputNumber < 1) {
    resultElement.innerHTML = "Please enter a positive number."
    return
  }

  var pyramid = ""
  for (var i = 1; i <= inputNumber; i++) {
    var spaces = " ".repeat(inputNumber - i)
    var asterisks = "*".repeat(i * 2 - 1)
    pyramid += spaces + asterisks + "<br>"
  }

  document.getElementById("result").innerHTML = pyramid
}
