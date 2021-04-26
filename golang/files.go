//check file exists
if _, err := os.Stat("/path/to/whatever"); os.IsNotExist(err) {
  // path/to/whatever does not exist
}
//handle err
