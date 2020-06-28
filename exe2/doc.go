/* Package exe2 splits templating/env expansion and command running into their own method calls.

This allows using a test runner which simply logs executions instead of returning something. He is useful if there are no outputs or changes following from the command.

The order of executing the template, splitting the command into args
and expanding the env variables may be important to you. If so simply
switch the TemplateSplitExpand method with another implementation.
It is only important that it returns exe2.SplitResult{Name:Args:Err:}.
*/
package exe2
