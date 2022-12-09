package day07

import shared.AotUtil

class Main {
    fun main() {
        val lines = AotUtil(this.javaClass.packageName).readLinesProduction()

        println("Part One: " + runSolutionPartOne(lines))
        println("Part Two: " + runSolutionPartTwo(lines))
    }

    class File(val size: Int = 0)
    class Directory(
        var dirMap: MutableMap<String, Directory>,
        var name: String,
        var parent: Directory?,
        private var dirChildren: MutableList<Directory> = mutableListOf<Directory>(),
        private var fileChildren: MutableList<File> = mutableListOf<File>()
    ) {
        init {
            dirMap.set(name, this)
        }

        fun calculateSize(): Int {
            var sum = 0
            for (child in this.fileChildren) {
                sum += child.size
            }
            for (child in this.dirChildren) {
                sum += child.calculateSize()
            }
            return sum
        }

        fun addDirectory(name: String) {
            val dir = Directory(this.dirMap, this.name + name, this)
            this.dirChildren.add(dir)
        }

        fun addFile(name: String, size: Int) {
            val file = File(size)
            this.fileChildren.add(file)
        }
    }

    private val dirMap = mutableMapOf<String, Directory>()

    private fun parseCommands(lines: List<String>): Directory {
        var currentDir: Directory? = null
        val root = Directory(dirMap, "root", null)
        for (line in lines) {
            when {
                line == "$ cd /" -> {
                    currentDir = root
                    continue
                }

                line.startsWith("dir") -> {
                    currentDir?.addDirectory(line.split(" ")[1])
                }

                Regex("^[0-9]+.*").matches(line) -> {
                    val (size, name) = line.split(" ")
                    currentDir?.addFile(name, size.toInt())
                }

                line == "\$ cd .." -> {
                    currentDir = currentDir?.parent
                }

                line.startsWith("$ cd") -> {
                    currentDir = dirMap[currentDir?.name + line.split(" ")[2]]
                }
            }

        }

        return root
    }

    fun runSolutionPartOne(lines: List<String>): Int {
        val root = parseCommands(lines)

        val listOfSizesLessThan100000 = mutableListOf<Int>()
        for ((_, dir) in root.dirMap) {
            val size = dir.calculateSize()
            if (size <= 100000) {
                listOfSizesLessThan100000.add(size)
            }
        }
        return listOfSizesLessThan100000.sum()
    }

    fun runSolutionPartTwo(lines: List<String>): Int {
        val root = parseCommands(lines)

        val unused = 70000000 - root.calculateSize()
        val target = 30000000 - unused

        val listOfSizes = mutableListOf<Int>()
        for ((_, dir) in root.dirMap) {
            val size = dir.calculateSize()

            if (size > target) {
                listOfSizes.add(size)
            }
        }
        return listOfSizes.sorted()[0]
    }
}

fun main() {
    Main().main()
}