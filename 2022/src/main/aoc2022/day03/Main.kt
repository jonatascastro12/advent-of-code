package day03

import shared.AotUtil

class Main {
    fun main() {
        val lines = AotUtil(this.javaClass.packageName).readLinesProduction()

        println("Part One: " + runSolutionPartOne(lines))
        println("Part Two: " + runSolutionPartTwo(lines))
    }

    fun convertCharToCode(char: Char): Int {
        if (Regex("[a-z]").matches(char.toString()))
            return char.code - 96

        if (Regex("[A-Z]").matches(char.toString()))
            return char.code - 38

        return 0
    }

    fun runSolutionPartOne(lines: List<String>): Int {
        val priorities = mutableListOf<Int>()

        lines.forEach { line ->
            val partOne = line.substring(0, line.length / 2)
            val partTwo = line.substring(line.length / 2, line.length)

            for (it in partTwo) {
                if (it in partOne) {
                    priorities.add(convertCharToCode(it))
                    break
                }
            }
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
}

fun main() {
    Main().main()
}