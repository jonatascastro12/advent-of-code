package day03

import java.io.File
import java.nio.file.Paths
import kotlin.text.Regex

fun readLines(dayPath: String): List<String> {
    val path = Paths.get("").toAbsolutePath().toString()
    val fileName = "$path/src/main/kotlin/$dayPath"
    return File(fileName).readLines()
}

fun main() {
    val lines = readLines("day03/input.txt")

    println("Part One: " + runSolutionPartOne(lines));
    println("Part Two: " + runSolutionPartTwo(lines));
}

fun convertCharToCode(char: Char): Int {
    val azRegex = Regex("[a-z]")
    val AZRegex = Regex("[A-Z]")

    if (azRegex.matches(char.toString()))
        return char.code - 96


    if (AZRegex.matches(char.toString()))
        return char.code - 38

    return 0
}

fun runSolutionPartOne(lines: List<String>): Int {
    var priorities = mutableListOf<Int>()

    lines.forEach { line ->
        var currentIntersection = mutableSetOf<Int>()

        val partOne = line.substring(0, line.length / 2)
        val partTwo = line.substring(line.length / 2, line.length)

        partTwo.forEach {
            if (it in partOne) {
                currentIntersection.add(convertCharToCode(it))
            }
        }
        priorities.add(currentIntersection.sum())
    }

    return priorities.sum()
}

fun runSolutionPartTwo(lines: List<String>): Int {
    val priorities = mutableListOf<Int>()

    lines.forEachIndexed { i, _ ->
        if (i % 3 == 2) {
            val first = lines[i - 2]
            val second = lines[i - 1]
            val third = lines[i]

            for (it in third) {
                if (it in first && it in second) {
                    priorities.add(convertCharToCode(it))
                    break
                }
            }
        }
    }

    return priorities.sum()
}