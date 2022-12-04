package day04

import shared.AotUtil
import kotlin.math.max
import kotlin.math.min

class Main {
    fun main() {
        val lines = AotUtil(this.javaClass.packageName).readLinesProduction()

        println("Part One: " + runSolutionPartOne(lines))
        println("Part Two: " + runSolutionPartTwo(lines))
    }

    fun runSolutionPartOne(lines: List<String>): Int {
        var count = 0

        lines.forEach { line ->
            val (f, s) = parseLine(line)

            val firstInSecond = f[0] >= s[0] && f[1] <= s[1]
            val secondInFirst = s[0] >= f[0] && s[1] <= f[1]

            if (firstInSecond || secondInFirst) count++
        }

        return count
    }

    fun runSolutionPartTwo(lines: List<String>): Int {
        var count = 0

        lines.forEach { line ->
            val (f, s) = parseLine(line)

            if (max(f[0], s[0]) <= min(f[1], s[1])) {
                count++
            }
        }

        return count
    }

    private fun parseLine(line: String): Pair<List<Int>, List<Int>> {
        val split = line.split(",")
        val f = split[0].split("-").map { i -> i.toInt() }
        val s = split[1].split("-").map { i -> i.toInt() }
        return Pair(f, s)
    }
}

fun main() {
    Main().main()
}