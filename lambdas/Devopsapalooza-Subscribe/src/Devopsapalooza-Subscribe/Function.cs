using Amazon.Lambda.Core;
using Amazon.DynamoDBv2;
using Amazon.DynamoDBv2.DocumentModel;

// Assembly attribute to enable the Lambda function's JSON input to be converted into a .NET class.
[assembly: LambdaSerializer(typeof(Amazon.Lambda.Serialization.SystemTextJson.DefaultLambdaJsonSerializer))]

namespace Devopsapalooza_Subscribe;

public class Function
{
    public async Task<bool> FunctionHandler(String input, ILambdaContext context)
    {
        var result = await SaveCustomer(input);

        return result;
    }

    private static AmazonDynamoDBClient client = new AmazonDynamoDBClient();
    private static string tableName = "Customers";

    private async Task<bool> SaveCustomer(String email)
    {
        try
        {
            var customersTable = Table.LoadTable(client, tableName);
            
            var customer = new Document();
            customer["Email"] = email;
            customer["Message"] = "Source: Devopsapalooza-Subscribe";
            customer["Date"] = DateTime.Now;

            await customersTable.PutItemAsync(customer);
        }
        catch (Exception e) 
        { 
            Console.WriteLine(e.Message); 
            return false;
        }
        return true;
    }
}
