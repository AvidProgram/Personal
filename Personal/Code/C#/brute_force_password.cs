using System;
using System.Text;
using System.Diagnostics;
using System.Numerics;
 
namespace BruteForcePasswordGeneration
{
    class Program
    {
        static void Main(string[] args)
        {
 
 
            char[] chars = { 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'};
 
 
            int passwordLength=0;
 
            Console.WriteLine("Enter the password length");
 
            passwordLength = Convert.ToInt32(Console.ReadLine());
 
            BigInteger iPossibilities = (BigInteger)Math.Pow((double)chars.Length, (double)passwordLength);
            Console.WriteLine("{0} words total. Press enter to continue;", iPossibilities);
            Console.ReadLine();
 
            var watch = Stopwatch.StartNew();
            for (BigInteger i = 0; i < iPossibilities; i++)
            {
                string theword = "";
                BigInteger val = i;
                for (int j = 0; j < passwordLength; j++)
                {
                    BigInteger ch = val % chars.Length;
                    theword = chars[(int)ch] + theword;
                    val = val / chars.Length;
                }
                 Console.WriteLine(theword);
            }
            watch.Stop();
            var elapsedMs = watch.ElapsedMilliseconds;
            Console.WriteLine("It took {0} seconds to generate {1} possible combinations", elapsedMs / 1000, iPossibilities);
            Console.ReadLine();
 
        }
    }
}