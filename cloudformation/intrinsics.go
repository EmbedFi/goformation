package cloudformation

import (
	"encoding/base64"
	"strings"
)

// Note: Intrinsic objects are Base64 encoded, to prevent escaping (backslash) issues
// with nested intrinsic functions.

// Ref creates a CloudFormation Reference to another resource in the template
func Ref(logicalName string) string {
	return encode(`{ "Ref": "` + logicalName + `" }`)
}

// GetAtt returns the value of an attribute from a resource in the template.
func GetAtt(logicalName string, attribute string) string {
	return encode(`{ "Fn::GetAtt": [ "` + logicalName + `", "` + attribute + `" ] }`)
}

// ImportValue returns the value of an output exported by another stack. You typically use this function to create cross-stack references. In the following example template snippets, Stack A exports VPC security group values and Stack B imports them.
func ImportValue(name string) string {
	return encode(`{ "Fn::ImportValue": "` + name + `" }`)
}

// Base64 returns the Base64 representation of the input string. This function is typically used to pass encoded data to Amazon EC2 instances by way of the UserData property
func Base64(input string) string {
	return encode(`{ "Fn::Base64": "` + input + `" }`)
}

// CIDR returns an array of CIDR address blocks. The number of CIDR blocks returned is dependent on the count parameter.
func CIDR(ipBlock, count, cidrBits string) string {
	return encode(`{ "Fn::Cidr" : [ "` + ipBlock + `", "` + count + `", "` + cidrBits + `" ] }`)
}

// FindInMap returns the value corresponding to keys in a two-level map that is declared in the Mappings section.
func FindInMap(mapName, topLevelKey, secondLevelKey string) string {
	return encode(`{ "Fn::FindInMap" : [ "` + mapName + `", "` + topLevelKey + `", "` + secondLevelKey + `" ] }`)
}

// GetAZs returns an array that lists Availability Zones for a specified region. Because customers have access to different Availability Zones, the intrinsic function Fn::GetAZs enables template authors to write templates that adapt to the calling user's access. That way you don't have to hard-code a full list of Availability Zones for a specified region.
func GetAZs(region string) string {
	return encode(`{ "Fn::GetAZs": "` + region + `" }`)
}

// Join appends a set of values into a single value, separated by the specified delimiter. If a delimiter is the empty string, the set of values are concatenated with no delimiter.
func Join(delimiter string, values []string) string {
	return encode(`{ "Fn::Join": [ "` + delimiter + `", [ "` + strings.Trim(strings.Join(values, `", "`), `, "`) + `" ] ] }`)
}

// Select returns a single object from a list of objects by index.
func Select(index string, list []string) string {
	return encode(`{ "Fn::Select": [ "` + index + `", [ "` + strings.Trim(strings.Join(list, `", "`), `, "`) + `" ] ] }`)
}

// Split splits a string into a list of string values so that you can select an element from the resulting string list, use the Fn::Split intrinsic function. Specify the location of splits with a delimiter, such as , (a comma). After you split a string, use the Fn::Select function to pick a specific element.
func Split(delimiter, source string) string {
	return encode(`{ "Fn::Split" : [ "` + delimiter + `", "` + source + `" ] }`)
}

// Sub substitutes variables in an input string with values that you specify. In your templates, you can use this function to construct commands or outputs that include values that aren't available until you create or update a stack.
func Sub(value string) string {
	return encode(`{ "Fn::Sub" : "` + value + `" }`)
}

// encode takes a string representation of an intrinsic function, and base64 encodes it.
// This prevents the escaping issues when nesting multiple layers of intrinsic functions.
func encode(value string) string {
	return base64.StdEncoding.EncodeToString([]byte(value))
}
