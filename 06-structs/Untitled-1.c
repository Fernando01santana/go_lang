nt PinoVelocidade = 3; //Ligado ao pino 1 do L293D  
int PinoVelocidade2 = 8; //Ligado ao pino 1 do L293D  
int Entrada1 = 2; //Ligado ao pino 2 do L293D  
int Entrada2 = 7; //Ligado ao pino 7 do L293D 
int Entrada3 = 4; //Ligado ao pino 2 do L293D  
int Entrada4 = 9; //Ligado ao pino 7 do L293D  
//Definição dos pinos dos sensores
#define pin_S1 5
#define pin_S2 6
bool Sensor1 = 0;
bool Sensor2 = 0;

//variável responsável por controlar a velocidade dos motores
int velocidade = 150;
void setup(){
//Setamos os pinos de controle dos motores como saída
  pinMode(3, OUTPUT);  
  pinMode(2, OUTPUT);  
  pinMode(7, OUTPUT); 
   
   pinMode(4, OUTPUT);  
  pinMode(8, OUTPUT);  
  pinMode(9, OUTPUT); 


  
//Setamos a direção inicial do motor como 0, isso fará com que ambos os motores girem para frente
digitalWrite(Entrada1, 0);
digitalWrite(Entrada2, 0);
digitalWrite(4, 0);
digitalWrite(9, 0);
//Setamos os pinos dos sensores como entrada
pinMode(pin_S1, INPUT);
pinMode(pin_S2, INPUT);
}
void loop(){
//Neste processo armazenamos o valor lido pelo sensor na variável que armazena tais dados.
Sensor1 = digitalRead(pin_S1);
Sensor2 = digitalRead(pin_S2);

//Aqui está toda a lógica de comportamento do robô: Para a cor branca atribuímos o valor 0 e, para a cor preta, o valor 1.
if((Sensor1 == 0) && (Sensor2 == 0)){ // Se detectar na extremidade das faixas duas cores brancas
analogWrite(3,velocidade); // Ambos motores ligam na mesma velocidade
analogWrite(8, velocidade);
}
if((Sensor1 == 1) && (Sensor2 == 0)){ // Se detectar um lado preto e o outro branco
analogWrite(2, 0); // O motor 1 desliga
analogWrite(9, velocidade); // O motor 2 fica ligado, fazendo assim o carrinho virar
}
if((Sensor1 == 0) && (Sensor2 == 1)){ // Se detectar um lado branco e o outro preto
analogWrite(2, velocidade); //O motor 1 fica ligado
analogWrite(9, 0); // O motor 2 desliga, fazendo assim o carrinho virar no outro sentido
}

}