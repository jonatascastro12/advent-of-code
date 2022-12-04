package day01

import shared.AotUtil

class Main {
    fun main() {
        val lines = AotUtil(this.javaClass.packageName).readLinesProduction()

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
}