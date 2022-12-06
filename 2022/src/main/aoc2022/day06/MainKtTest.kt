package day06

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class MainKtTestDay06 {
    lateinit var lines: List<String>

    @Test
    fun countUntilMarker() {
        assertEquals(7, Main().countUntilMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
        assertEquals(5, Main().countUntilMarker("bvwbjplbgvbhsrlpgdmjqwftvncz"))
        assertEquals(6, Main().countUntilMarker("nppdvjthqldpwncqszvftbrmjlhg"))
        assertEquals(10, Main().countUntilMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
        assertEquals(11, Main().countUntilMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
    }

    @Test
    fun countUntilMarkerFourteen() {
        assertEquals(19, Main().countUntilMarker("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14))
        assertEquals(23, Main().countUntilMarker("bvwbjplbgvbhsrlpgdmjqwftvncz", 14))
        assertEquals(23, Main().countUntilMarker("nppdvjthqldpwncqszvftbrmjlhg", 14))
        assertEquals(29, Main().countUntilMarker("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14))
        assertEquals(26, Main().countUntilMarker("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14))
    }
}