package shared

import org.junit.jupiter.api.Assertions.assertEquals
import org.junit.jupiter.api.Test

class AotUtilTest {
    var aotUtil: AotUtil = AotUtil("shared")

    @Test
    fun readLinesProduction() {
        assertEquals(aotUtil.readLinesProduction(), listOf("1", "2", "3", "4", "5"))
    }

    @Test
    fun readLinesTest() {
        assertEquals(aotUtil.readLinesTest(), listOf("a", "b", "c", "d", "e"))
    }
}