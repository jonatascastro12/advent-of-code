package day06

import shared.AotUtil

class Main {
    fun main() {
        val lines = AotUtil(this.javaClass.packageName).readLinesProduction()

        println("Part One: " + runSolutionPartOne(lines))
        println("Part Two: " + runSolutionPartTwo(lines))
    }

    fun runSolutionPartOne(lines: List<String>): Int {
        return countUntilMarker(lines[0])
    }

    fun runSolutionPartTwo(lines: List<String>): Int {
        return countUntilMarker(lines[0], 14)
    }

    fun countUntilMarker(signal: String, markerDistinctCount: Int = 4): Int {
        var i = markerDistinctCount - 1
        while (i < signal.length) {
            val substr = signal.substring(i - (markerDistinctCount - 1), i + 1).toCharArray()

            if (substr.toSet().size == markerDistinctCount) {
                return i + 1
            }
            i++
        }
        return -1
    }
}

fun main() {
    Main().main()
}