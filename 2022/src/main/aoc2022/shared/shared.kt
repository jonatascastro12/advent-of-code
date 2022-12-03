package shared

import java.io.File
import java.nio.file.Paths

fun readLines(dayPath: String): List<String> {
    val path = Paths.get("").toAbsolutePath().toString()
    val fileName = "$path/src/main/aoc2022/$dayPath"
    return File(fileName).readLines()
}
