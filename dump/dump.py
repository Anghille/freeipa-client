import urllib.request
import sys
import re
import inspect
import json
import types
import os


ERRORS_PY_URL = "https://raw.githubusercontent.com/freeipa/freeipa/ipa-4-10/ipalib/errors.py"

def should_keep(l, import_regex):
    """Check whether a line is not an import line"""
    return (import_regex.match(l) is None)

def main(): 

    # Regex to match import lines
    # This regex is composed of two main groups, each enclosed in parentheses.

    # 1. ^(from [\w\.]+ )? - This is the first group.
    #    ^ means the start of the line.
    #    (from [\w\.]+ ) is matching the "from" keyword followed by any word character or dot one or more times, followed by a space.
    #    ? after this group means that this group is optional - it matches "from x" or nothing at all.
    
    # 2. import \w+( as \w+)?$ - This is the second group.
    #    import \w+ matches the "import" keyword followed by any word character one or more times.
    #    ( as \w+)? is another group, following the same pattern as the first one - it matches " as y" or nothing at all.
    #    $ means the end of the line.
    
    # The regular expression, therefore, matches both "import x" and "from x import y" formats, with optional "as z" at the end.
    import_regex = re.compile(r"^(from [\w\.]+ )?import \w+( as \w+)?$")

    # Fetch the errors.py script
    errors_py_str = urllib.request.urlopen(ERRORS_PY_URL).read().decode('utf-8')
    
    # Filter out import lines
    errors_py_str = "\n".join(
        [l for l in errors_py_str.splitlines() if should_keep(l, import_regex)])
    
    # Add dummy classes and objects to make the script run correctly
    errors_py_str = """
class Six:
    PY3 = True
six = Six()
ungettext = None
class Messages:
    def iter_messages(*args):
        return []
messages = Messages()
    """ + errors_py_str
    
    # Create a new module and execute the fetched script in its context
    errors_mod = types.ModuleType('errors')
    exec(errors_py_str, errors_mod.__dict__)
    
    # Extract error codes from the module
    error_codes = [
        {
            "name": k,
            "errno": v.errno
        } for k, v in inspect.getmembers(errors_mod)
        if hasattr(v, '__dict__') and type(v.__dict__.get('errno', None)) == int
    ]

    # Sort error codes
    error_codes.sort(key=lambda x: x["errno"])
    
    # Find the current directory of the script and go up two levels to the root of the project
    project_root = os.path.dirname(os.path.dirname(os.path.abspath(__file__)))
    output_file_path = os.path.join(project_root, 'data/errors.json')
    
    # Write error codes to file
    with open(output_file_path, 'w') as f:
        json.dump(error_codes, f)


if __name__=="__main__":
    main()
