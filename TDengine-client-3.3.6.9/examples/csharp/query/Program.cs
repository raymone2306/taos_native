using System.Text;
using TDengine.Driver;
using TDengine.Driver.Client;

namespace TDengineExample
{
    internal class QueryExample
    {
        public static void Main(string[] args)
        {
            var builder = new ConnectionStringBuilder("host=localhost;port=6030;username=root;password=taosdata");
            using (var client = DbDriver.Open(builder))
            {
                try
                {
                    string query = "SELECT * FROM power.meters";
                    using (var rows = client.Query(query))
                    {
                        while (rows.Read())
                        {
                            Console.WriteLine(
                                $"{((DateTime)rows.GetValue(0)):yyyy-MM-dd HH:mm:ss.fff}, {rows.GetValue(1)}, {rows.GetValue(2)}, {rows.GetValue(3)}, {rows.GetValue(4)}, {Encoding.UTF8.GetString((byte[])rows.GetValue(5))}");
                        }
                    }
                }
                catch (Exception e)
                {
                    Console.WriteLine(e.ToString());
                    throw;
                }
            }
        }
    }
}