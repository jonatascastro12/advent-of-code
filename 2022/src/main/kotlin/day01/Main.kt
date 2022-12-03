package day01

import java.io.File
import java.nio.file.Paths

fun main() {
    val path = Paths.get("").toAbsolutePath().toString()
    val fileName = "$path/src/main/kotlin/Day_01/input.txt"
    val lines = File(fileName).readLines()
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