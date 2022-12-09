package day07

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.BeforeEach
import org.junit.jupiter.api.Test
import shared.AotUtil

class MainKtTestDay07 {
    lateinit var lines: List<String>

    @BeforeEach
    fun setUp() {
        this.lines = AotUtil(this.javaClass.packageName).readLinesTest()
    }

    @Test
    fun runSolutionPartOne() {
        assertEquals(Main().runSolutionPartOne(this.lines), 95437)
    }

    @Test
    fun runSolutionPartTwo() {
        assertEquals(Main().runSolutionPartTwo(this.lines), 24933642)
    }
}