package day03

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.BeforeEach
import org.junit.jupiter.api.Test
import shared.readLines

class MainKtTest {
    lateinit var lines: List<String>

    @BeforeEach
    fun setUp() {
        this.lines = readLines("day03/input_test.txt")
    }

    @Test
    fun runSolutionPartOne() {
        assertEquals(runSolutionPartOne(this.lines), 157)
    }

    @Test
    fun runSolutionPartTwo() {
        assertEquals(runSolutionPartTwo(this.lines), 70)
    }

    @Test
    fun convertCharToCode() {
        assertEquals(convertCharToCode('a'), 1)
        assertEquals(convertCharToCode('b'), 2)
        assertEquals(convertCharToCode('p'), 16)
        assertEquals(convertCharToCode('z'), 26)
        assertEquals(convertCharToCode('A'), 27)
        assertEquals(convertCharToCode('B'), 28)
        assertEquals(convertCharToCode('Z'), 52)
    }
}