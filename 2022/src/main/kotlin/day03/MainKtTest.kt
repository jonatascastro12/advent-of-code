package day03

import org.junit.jupiter.api.Assertions.*
import org.junit.jupiter.api.BeforeEach
import org.junit.jupiter.api.Test

class MainKtTest {
    lateinit var lines: List<String>

    @BeforeEach
    fun setUp() {
        this.lines = readLines("day03/input_test.txt")
    }

    @Test
    fun runSolutionPartOne() {
        assertEquals(day03.runSolutionPartOne(this.lines), 157)
    }

    @Test
    fun runSolutionPartTwo() {
        assertEquals(day03.runSolutionPartTwo(this.lines), 70)
    }

    @Test
    fun convertCharToCode() {
        assertEquals(day03.convertCharToCode('a'), 1)
        assertEquals(day03.convertCharToCode('b'), 2)
        assertEquals(day03.convertCharToCode('p'), 16)
        assertEquals(day03.convertCharToCode('z'), 26)
        assertEquals(day03.convertCharToCode('A'), 27)
        assertEquals(day03.convertCharToCode('B'), 28)
        assertEquals(day03.convertCharToCode('Z'), 52)
    }
}