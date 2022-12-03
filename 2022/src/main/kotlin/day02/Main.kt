package day02

import java.io.File
import java.nio.file.Paths

fun main() {
    val path = Paths.get("").toAbsolutePath().toString()
    val fileName = "$path/src/main/kotlin/Day_02/input.txt"
    val lines = File(fileName).readLines()

    /**
     * A -> Rock
     * B -> Paper
     * C -> Scissors
     *
     * X -> Rock 1
     * Y -> Paper 2
     * Z -> Scissors 3
     *
     * win = 6
     * draw = 3
     * lost = 0
     */
    val rulePart1 = mapOf(
        "A X" to 1 + 3,
        "A Y" to 2 + 6,
        "A Z" to 3 + 0,
        "B X" to 1 + 0,
        "B Y" to 2 + 3,
        "B Z" to 3 + 6,
        "C X" to 1 + 6,
        "C Y" to 2 + 0,
        "C Z" to 3 + 3,
    )


    /**
     * A -> Rock 1
     * B -> Paper 2
     * C -> Scissors 3
     *
     * X -> Lose
     * Y -> Draw
     * Z -> Win
     *
     */
    val rulePart2 = mapOf(
        "A X" to 0 + 3,
        "A Y" to 3 + 1,
        "A Z" to 6 + 2,
        "B X" to 0 + 1,
        "B Y" to 3 + 2,
        "B Z" to 6 + 3,
        "C X" to 0 + 2,
        "C Y" to 3 + 3,
        "C Z" to 6 + 1,
    )

    var totalScorePart1 = 0;
    var totalScorePart2 = 0;

    lines.forEach {
        totalScorePart1 += rulePart1[it]!!;
        totalScorePart2 += rulePart2[it]!!;
    }

    println("First Part: " + totalScorePart1)
    println(
        "Second Part: " + totalScorePart2
    )
}