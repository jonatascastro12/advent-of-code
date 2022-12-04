package day03

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.BeforeEach
import org.junit.jupiter.api.Test
import shared.AotUtil

class MainKtTestDay03 {
    lateinit var lines: List<String>

    @BeforeEach
    fun setUp() {
        this.lines = AotUtil(this.javaClass.packageName).readLinesTest()
    }

    @Test
    fun runSolutionPartOne() {
        assertEquals(Main().runSolutionPartOne(this.lines), 157)
    }

    @Test
    fun runSolutionPartTwo() {
        assertEquals(Main().runSolutionPartTwo(this.lines), 70)
    }

    @Test
    fun convertCharToCode() {
        assertEquals(Main().convertCharToCode('a'), 1)
        assertEquals(Main().convertCharToCode('b'), 2)
        assertEquals(Main().convertCharToCode('p'), 16)
        assertEquals(Main().convertCharToCode('z'), 26)
        assertEquals(Main().convertCharToCode('A'), 27)
        assertEquals(Main().convertCharToCode('B'), 28)
        assertEquals(Main().convertCharToCode('Z'), 52)
    }
}