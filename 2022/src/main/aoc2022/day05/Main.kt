package day05

import shared.AotUtil

class Main {
    fun main() {
        val lines = AotUtil(this.javaClass.packageName).readLinesProduction()

        println("Part One: " + runSolutionPartOne(lines))
        println("Part Two: " + runSolutionPartTwo(lines))
    }

    fun runSolutionPartOne(lines: List<String>): Int {
        return 0
    }

    fun runSolutionPartTwo(lines: List<String>): Int {
        return 0
    }
}

fun main() {
    Main().main()
}