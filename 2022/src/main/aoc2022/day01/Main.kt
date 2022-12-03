package day01

import shared.readLines
import java.io.File
import java.nio.file.Paths

fun main() {
    val lines = readLines("day01/input.txt")

    val elfCalories = mutableListOf<Int>()

    var lastElf = 0
    lines.forEach {
        if (it != "") lastElf += it.toInt()
        else {
            elfCalories.add(lastElf)
            lastElf = 0
        }
    }
    println("First Part: " + elfCalories.sortedDescending()[0])
    println(
        "Second Part: " + (
                elfCalories.sortedDescending()[0] +
                        elfCalories.sortedDescending()[1] +
                        elfCalories.sortedDescending()[2])
    )
}