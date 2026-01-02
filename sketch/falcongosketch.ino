#include <dht11.h>
dht11 DHT;
#define DHT11_PIN_00 12
#define DHT11_PIN_01 11

#define SWITCH_PIN_00 3
#define SWITCH_PIN_01 4
#define SWITCH_PIN_02 5
#define SWITCH_PIN_03 6
#define SWITCH_PIN_04 7
#define SWITCH_PIN_05 8
#define SWITCH_PIN_06 9
#define SWITCH_PIN_07 10

#define ACS712_PIN_00 0
#define LIGHT_PIN_00 1

#define CDELAY 100

void setup() {
	pinMode(DHT11_PIN_00, INPUT);
	pinMode(DHT11_PIN_01, INPUT);

	pinMode(SWITCH_PIN_00, OUTPUT);
	pinMode(SWITCH_PIN_01, OUTPUT);
	pinMode(SWITCH_PIN_02, OUTPUT);
	pinMode(SWITCH_PIN_03, OUTPUT);
	pinMode(SWITCH_PIN_04, OUTPUT);
	pinMode(SWITCH_PIN_05, OUTPUT);
	pinMode(SWITCH_PIN_06, OUTPUT);
	pinMode(SWITCH_PIN_07, OUTPUT);
        
	Serial.begin(9600);
        Serial.println("Ready");
	//Serial.println("Type,\tstatus,\tHumidity (%),\tTemperature (C)");
}

int readline(int readch, char *buffer, int len) {
	static int pos = 0;
	int rpos;

	if (readch > 0) {
		switch (readch) {
		case '\n':
			break;
		case '\r':
			rpos = pos;
			pos = 0;
			return rpos;
		default:
			if (pos < len-1) {
				buffer[pos++] = readch;
				buffer[pos] = 0;
			}
		}
	}
	return -1;
}

int getMaxValue(int inputpin) {
	int sensorValue;
	int sensorMax = 0;
	uint32_t start_time = millis();
	while((millis()-start_time) < 1000) {
		sensorValue = analogRead(inputpin);
		if (sensorValue > sensorMax) {
			sensorMax = sensorValue;
		}
	}
	return sensorMax;
}

void loop() {
	static char buffer[80];

	static char switches[8];

        if (readline(Serial.read(), buffer, 80) > 0) {
		if (strcmp(buffer,"get temp00")==0) {
	                int ok=0;
	                
			int chk = DHT.read(DHT11_PIN_00);
			switch (chk){
			case DHTLIB_OK:
				ok=1;
				break;
			case DHTLIB_ERROR_CHECKSUM: 
				//Serial.println("Checksum error"); 
				break;
			case DHTLIB_ERROR_TIMEOUT: 
				//Serial.println("Time out error"); 
				break;
			default: 
				//Serial.println("Unknown error"); 
				break;
			}
	
			if (ok == 1) {
	                	Serial.write("temp00:");
				Serial.println(DHT.temperature,1);
				delay(CDELAY);
			}
		}

		if (strcmp(buffer,"get temp01")==0) {
	                int ok=0;
	                
			int chk = DHT.read(DHT11_PIN_01);
			switch (chk){
			case DHTLIB_OK:
				ok=1;
				break;
			case DHTLIB_ERROR_CHECKSUM: 
				//Serial.println("Checksum error"); 
				break;
			case DHTLIB_ERROR_TIMEOUT: 
				//Serial.println("Time out error"); 
				break;
			default: 
				//Serial.println("Unknown error"); 
				break;
			}
	
			if (ok == 1) {
	                	Serial.write("temp01:");
				Serial.println(DHT.temperature,1);
				delay(CDELAY);
			}
		}
	
		if (strcmp(buffer,"get hum00")==0) {
	                int ok=0;
	                
			int chk = DHT.read(DHT11_PIN_00);
			switch (chk){
			case DHTLIB_OK:
				ok=1;
				break;
			case DHTLIB_ERROR_CHECKSUM: 
				//Serial.println("Checksum error"); 
				break;
			case DHTLIB_ERROR_TIMEOUT: 
				//Serial.println("Time out error"); 
				break;
			default: 
				//Serial.println("Unknown error"); 
				break;
			}
	
			if (ok == 1) {
	                	Serial.write("hum00:");
				Serial.println(DHT.humidity,1);
				delay(CDELAY);
			}
		}

		if (strcmp(buffer,"get hum01")==0) {
	                int ok=0;
	                
			int chk = DHT.read(DHT11_PIN_01);
			switch (chk){
			case DHTLIB_OK:
				ok=1;
				break;
			case DHTLIB_ERROR_CHECKSUM: 
				//Serial.println("Checksum error"); 
				break;
			case DHTLIB_ERROR_TIMEOUT: 
				//Serial.println("Time out error"); 
				break;
			default: 
				//Serial.println("Unknown error"); 
				break;
			}
	
			if (ok == 1) {
	                	Serial.write("hum01:");
				Serial.println(DHT.humidity,1);
				delay(CDELAY);
			}
		}
		if (strcmp(buffer,"get light00")==0) {
			float light;
			light = analogRead(LIGHT_PIN_00);
			Serial.write("light00:");
			Serial.println(light);
			delay(CDELAY);
		}
		if (strcmp(buffer,"get dc00")==0) {
			float average = 0;
			for(int i = 0; i < 1000; i++) {
				average = average +(.044 * analogRead(ACS712_PIN_00) - 3.78) / 1000;
				delay(1);
			}
			Serial.write("dc00:");
			Serial.println(average);
			delay(CDELAY);
		}
		if (strcmp(buffer,"get ac00")==0) {
			int sensor_max;
			float amplitude_current;
			float effective_value;
			sensor_max = getMaxValue(ACS712_PIN_00);
			amplitude_current=(float) (.044 * sensor_max - 3.78);
			effective_value=amplitude_current/1.414;
			Serial.write("ac00:");
			Serial.println(effective_value);
			delay(CDELAY);
		}
		if (strcmp(buffer,"get switch00")==0) {
			Serial.write("switch00:");
			if (switches[0] == 0) {
				Serial.println("0");
			} else {
				Serial.println("1");
			}
			delay(CDELAY);
		}
		if (strcmp(buffer,"get switch01")==0) {
			Serial.write("switch01:");
			if (switches[1] == 0) {
				Serial.println("0");
			} else {
				Serial.println("1");
			}
			delay(CDELAY);
		}
		if (strcmp(buffer,"get switch02")==0) {
			Serial.write("switch02:");
			if (switches[2] == 0) {
				Serial.println("0");
			} else {
				Serial.println("1");
			}
			delay(CDELAY);
		}
		if (strcmp(buffer,"get switch03")==0) {
			Serial.write("switch03:");
			if (switches[3] == 0) {
				Serial.println("0");
			} else {
				Serial.println("1");
			}
			delay(CDELAY);
		}
		if (strcmp(buffer,"get switch04")==0) {
			Serial.write("switch04:");
			if (switches[4] == 0) {
				Serial.println("0");
			} else {
				Serial.println("1");
			}
			delay(CDELAY);
		}
		if (strcmp(buffer,"get switch05")==0) {
			Serial.write("switch05:");
			if (switches[5] == 0) {
				Serial.println("0");
			} else {
				Serial.println("1");
			}
			delay(CDELAY);
		}
		if (strcmp(buffer,"get switch06")==0) {
			Serial.write("switch06:");
			if (switches[6] == 0) {
				Serial.println("0");
			} else {
				Serial.println("1");
			}
			delay(CDELAY);
		}
		if (strcmp(buffer,"get switch07")==0) {
			Serial.write("switch07:");
			if (switches[7] == 0) {
				Serial.println("0");
			} else {
				Serial.println("1");
			}
			delay(CDELAY);
		}
		if (strcmp(buffer,"on switch00")==0) {
			digitalWrite(SWITCH_PIN_00, HIGH);
			switches[0]=1;
			delay(CDELAY);
		}
		if (strcmp(buffer,"off switch00")==0) {
			digitalWrite(SWITCH_PIN_00, LOW);
			switches[0]=0;
			delay(CDELAY);
		}
		if (strcmp(buffer,"on switch01")==0) {
			digitalWrite(SWITCH_PIN_01, HIGH);
			switches[1]=1;
			delay(CDELAY);
		}
		if (strcmp(buffer,"off switch01")==0) {
			digitalWrite(SWITCH_PIN_01, LOW);
			switches[1]=0;
			delay(CDELAY);
		}
		if (strcmp(buffer,"on switch02")==0) {
			digitalWrite(SWITCH_PIN_02, HIGH);
			switches[2]=1;
			delay(CDELAY);
		}
		if (strcmp(buffer,"off switch02")==0) {
			digitalWrite(SWITCH_PIN_02, LOW);
			switches[2]=0;
			delay(CDELAY);
		}
		if (strcmp(buffer,"on switch03")==0) {
			digitalWrite(SWITCH_PIN_03, HIGH);
			switches[3]=1;
			delay(CDELAY);
		}
		if (strcmp(buffer,"off switch03")==0) {
			digitalWrite(SWITCH_PIN_03, LOW);
			switches[3]=0;
			delay(CDELAY);
		}
		if (strcmp(buffer,"on switch04")==0) {
			digitalWrite(SWITCH_PIN_04, HIGH);
			switches[4]=1;
			delay(CDELAY);
		}
		if (strcmp(buffer,"off switch04")==0) {
			digitalWrite(SWITCH_PIN_04, LOW);
			switches[4]=0;
			delay(CDELAY);
		}
		if (strcmp(buffer,"on switch05")==0) {
			digitalWrite(SWITCH_PIN_05, HIGH);
			switches[5]=1;
			delay(CDELAY);
		}
		if (strcmp(buffer,"off switch05")==0) {
			digitalWrite(SWITCH_PIN_05, LOW);
			switches[5]=0;
			delay(CDELAY);
		}
		if (strcmp(buffer,"on switch06")==0) {
			digitalWrite(SWITCH_PIN_06, HIGH);
			switches[6]=1;
			delay(CDELAY);
		}
		if (strcmp(buffer,"off switch06")==0) {
			digitalWrite(SWITCH_PIN_06, LOW);
			switches[6]=0;
			delay(CDELAY);
		}
		if (strcmp(buffer,"on switch07")==0) {
			digitalWrite(SWITCH_PIN_07, HIGH);
			switches[7]=1;
			delay(CDELAY);
		}
		if (strcmp(buffer,"off switch07")==0) {
			digitalWrite(SWITCH_PIN_07, LOW);
			switches[7]=0;
			delay(CDELAY);
		}
	}
}
