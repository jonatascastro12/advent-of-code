package day05

import shared.AotUtil
import java.util.*

class Main {
    fun main() {
        val lines = AotUtil(this.javaClass.packageName).readLinesProduction()

        println("Part One: " + runSolutionPartOne(lines))
        println("Part Two: " + runSolutionPartTwo(lines))
    }

    private class Move(val count: Int?, val from: Int?, val to: Int?)

    private fun parseMoves(lines: List<String>, maxStackSize: Int): List<Move> {
        val moves = mutableListOf<Move>()
        var startMoves = maxStackSize + 3

        var newRawMoveLine: String = lines[startMoves]
        while (newRawMoveLine != "") {
            val moveRegex =
                Regex("move (?<count>\\d+) from (?<from>\\d+) to (?<to>\\d+)").matchEntire(newRawMoveLine)!!.groups as? MatchNamedGroupCollection

            val move = Move(
                count = moveRegex?.get("count")?.value?.toInt(),
                from = moveRegex?.get("from")?.value?.toInt(),
                to = moveRegex?.get("to")?.value?.toInt(),
            )

            moves.add(move)

            startMoves++
            newRawMoveLine = if (lines.size > startMoves) lines[startMoves] else ""
        }

        return moves
    }

    private fun parseStacks(lines: List<String>): Pair<List<Stack<Char>>, Int> {
        var i = 0
        while (lines[i][1] != '1') {
            i++
        }
        val maxStackSize = i - 1
        val numOfColumns = lines[i].split("   ").size

        val stacks = mutableListOf<Stack<Char>>()

        for (c in 1..numOfColumns) {
            val stack = Stack<Char>()
            for (row in maxStackSize downTo 0) {
                val col = ((c - 1) * 4) + 1
                try {
                    if (lines[row][col] != ' ') {
                        stack.push(lines[row][col])
                    }
                } catch (_: StringIndexOutOfBoundsException) {
                }
            }
            stacks.add(stack)
        }
        return stacks to maxStackSize
    }

    /*               1   5   9   13
                        [D]           0
                    [N] [C]           1
                    [Z] [M] [P] [T]   2
                     1   2   3   4
     */
    fun runSolutionPartOne(lines: List<String>): String {
        val (stacks, maxStackSize) = parseStacks(lines)
        val moves = parseMoves(lines, maxStackSize)


        for (move in moves) {
            for (i in 0 until move.count!!) {
                stacks[move.to!! - 1].push(stacks[move.from!! - 1].pop())
            }
        }

        return (stacks.indices).map { stacks[it].last() }.joinToString("")
    }

    fun runSolutionPartTwo(lines: List<String>): String {
        val (stacks, maxStackSize) = parseStacks(lines)
        val moves = parseMoves(lines, maxStackSize)


        for (move in moves) {
            val from = stacks[move.from!! - 1]
            val to = stacks[move.to!! - 1]
            val toMove = mutableListOf<Char>()

            for (i in 0 until move.count!!) {
                toMove.add(from.pop())
            }

            toMove.reversed().map { to.push(it) }
        }

        return (stacks.indices).map { stacks[it].last() }.joinToString("")
    }
}

fun main() {
    Main().main()
}