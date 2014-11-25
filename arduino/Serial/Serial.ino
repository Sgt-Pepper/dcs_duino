/* UART Example, any character received on either the real
 serial port, or USB serial (or emulated serial to the
 Arduino Serial Monitor when using non-serial USB types)
 is printed as a message to both ports.
 
 This example code is in the public domain.
 */

// set this to the hardware serial port you wish to use
#include <LiquidCrystal.h>


//define all possible commands
enum Color { 
  CSMP, RED, ORANGE, YELLOW, GREEN,  BLUE, PURPLE };

// initialize the library with the numbers of the interface pins
LiquidCrystal lcd(12, 11, 5, 4, 3, 2);
const int LED_PIN = 13;
const int LINE_BUFFER_SIZE = 80; // max line length is one less than this
boolean led = false;

void setup() {
  // set up the LCD's number of columns and rows: 
  lcd.begin(16, 2);
  Serial.begin(9600);
  pinMode(LED_PIN, OUTPUT); 


}

void loop() {
  Serial.println("Hello");
  led = (!led);
  if(led)
    digitalWrite(LED_PIN,HIGH);
  else
    digitalWrite(LED_PIN,LOW);


  // Read command

  char line[LINE_BUFFER_SIZE];
  if (read_line(line, sizeof(line)) < 0) {
    Serial.println("Error: line too long");
    return; // skip command processing and try again on next iteration of loop
  }

  String sline = String(line);
  // Process command
  int endCommand = sline.indexOf('=');
  if (endCommand <0) {
    Serial.println("no = found");
    return;
  }
  Serial.println("len");
  Serial.println(sline.length());
  Serial.println("ec+1");
  Serial.println(endCommand +1);
  if (sline.length() < endCommand +2) {
    Serial.println("no argument");
    return;
  }
  String command = sline.substring(0,endCommand);
  String argument = sline.substring(endCommand+1,sline.length());

  Serial.println("received the command ");
  Serial.println(command);
  Serial.println("received the command ");
  Serial.println(command);


  if (command == "CMSP"){
    doCSMP(argument);

  } 
  else {
    Serial.print("Error: unknown command: \"");
    Serial.print(command);

  }
}


int read_line(char* buffer, int bufsize){
  for (int index = 0; index < bufsize; index++) {
    // Wait until characters are available
    while (Serial.available() == 0) {
    }

    char ch = Serial.read(); // read next character
    //Serial.print(ch); // echo it back: useful with the serial monitor (optional)

    if (ch == '\n') {
      buffer[index] = 0; // end of line reached: null terminate string
      return index; // success: return length of string (zero if string is empty)
    }

    buffer[index] = ch; // Append character to buffer
  }

  // Reached end of buffer, but have not seen the end-of-line yet.
  // Discard the rest of the line (safer than returning a partial line).

  char ch;
  do {
    // Wait until characters are available
    while (Serial.available() == 0) {
    }
    ch = Serial.read(); // read next character (and discard it)
    // Serial.print(ch); // echo it back
  } 
  while (ch != '\n');
  buffer[0] = 0; // set buffer to empty string even though it should not be used
  return -1; // error: return negative one to indicate the input was too long
}


void doCSMP(String argument){
  if(argument.length() <32){
    if(argument.length() <16){
      setLcdText(argument,"");
    }
    else{
      setLcdText(argument.substring(0,16),argument.substring(16,argument.length()-1));   
    }
  }
  else{
    setLcdText(argument.substring(0,16),argument.substring(16,32));
  }

}

void setLcdText(String line1, String line2){
  lcd.clear();
  lcd.print(line1);
  lcd.setCursor(0,1);//second row
  lcd.print(line2);

}








