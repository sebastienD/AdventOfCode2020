using System;
using System.IO;
using System.Linq;

namespace day02
{
    public class Program
    {
        public static void Main(string[] args)
        {
            var count = 0;
            ;
            foreach (string line in File.ReadLines("../../../input.txt"))
            //foreach (var line in new string[]{"1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"})
            {
                if (isValidPasswordWithPosition(line))
                {
                    count++;
                }
            }
            Console.WriteLine($"There is {count} correct password.");
        }

        private static bool isValidPassword(string line)
        {
            string[] comp = line.Split(new char[] { ' ', '-', ':' }, StringSplitOptions.RemoveEmptyEntries);
            var min = int.Parse(comp[0]);
            var max = int.Parse(comp[1]);
            var letter = comp[2].ToCharArray()[0];
            var word = comp[3];
            //Console.WriteLine($"{min} {max} {letter} {word}");
            var result = word.Count(c => c == letter);
            return result >= min && result <= max;
        }
        
        private static bool isValidPasswordWithPosition(string line)
        {
            string[] comp = line.Split(new char[] { ' ', '-', ':' }, StringSplitOptions.RemoveEmptyEntries);
            var pos1 = int.Parse(comp[0])-1;
            var pos2 = int.Parse(comp[1])-1;
            var letter = comp[2].ToCharArray()[0];
            var word = comp[3];
            return (word[pos1] == letter && word[pos2] != letter) || (word[pos1] != letter && word[pos2] == letter);
        }
    }
}
