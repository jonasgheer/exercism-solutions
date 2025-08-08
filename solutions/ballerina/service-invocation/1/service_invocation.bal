import ballerina/http;
import ballerina/io;

final http:Client brainyQuoteClient = check new ("http://localhost:9095/brainyquote");

public function main() {
    string|http:ClientError quote = brainyQuoteClient->/;
    io:println(quote);
}
