using System;
using System.Collections.Generic;
using System.IO;
using System.Security.Policy;
using System.Text.RegularExpressions;

namespace day04
{
    public class Program
    {
        private static readonly Dictionary<string, Func<string, bool>> Checker = new Dictionary<string, Func<string, bool>>
        {
            { "byr", (value) => int.TryParse(value, out int val) && val >= 1920 && val <= 2002 },
            { "iyr", (value) => int.TryParse(value, out int val) && val >= 2010 && val <= 2020 },
            { "eyr", (value) => int.TryParse(value, out int val) && val >= 2020 && val <= 2030 },
            { "hgt", (value) => Regex.Match(value, "^(1([5-8]{1}[0-9]{1}|9[0-3]{1})cm|(59|6[0-9]|7[0-6])in)$").Success },
            { "hcl", (value) => Regex.Match(value, "^#[0-9a-f]{6}$").Success },
            { "ecl", (value) => Regex.Match(value, "^(amb|blu|brn|gry|grn|hzl|oth)$").Success },
            { "pid", (value) => Regex.Match(value, "^[0-9]{9}$").Success },
            { "cid", (value) => true}
        };
        
        public static void Main(string[] args)
        {
                var part = "";
                var count = 0;
                foreach (string line in File.ReadLines("../../../input.txt"))
                {
                    if (line == "")
                    {
                        if (ChekPass(part))
                        {
                            count++;
                        }

                        part = "";
                        continue;
                    }

                    part += " " + line;
                }

                if (part != "" && ChekPass(part))
                {
                    count++;
                }
                Console.WriteLine($"There is {count} correct pass.");
            }

        private static bool ChekPass(string part)
        {
            var values = part.Split(new char[] { ' ', ':' }, StringSplitOptions.RemoveEmptyEntries);
            var table = new HashSet<string>();
            for (int i = 0; i < values.Length; i+=2)
            {
                if (Checker[values[i]](values[i+1]))
                {
                    table.Add(values[i]);
                }
            }
            return table.Contains("byr") && table.Contains("iyr") && table.Contains("eyr") &&
                   table.Contains("hgt") && table.Contains("hcl") && table.Contains("ecl") &&
                   table.Contains("pid");
        }
    }
}