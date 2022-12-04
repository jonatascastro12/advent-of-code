package shared

import java.io.File
import java.nio.file.Paths

class AotUtil(private var dayPackageName: String) {
    private fun readLines(inputFileName: String = "input.txt"): List<String> {
        val path = Paths.get("").toAbsolutePath().toString()
        val fileName = "$path/src/main/aoc2022/${this.dayPackageName}/$inputFileName"
        return File(fileName).readLines()
    }

    fun readLinesProduction(): List<String> {
        return this.readLines()
    }

    fun readLinesTest(): List<String> {
        return this.readLines("input_test.txt")
    }
}
