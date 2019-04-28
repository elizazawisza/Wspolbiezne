package config is

   TasksListSize : constant Integer  := 20;
   ItemsSize : constant Integer := 20;

   TasksQuantity : constant Integer := 100;

   CeoDelay : constant Float := 0.30; 
   EmployerDelay : constant Float := 0.85;
   ClientDelay : constant Float := 1.3;

   EmployersAmount : constant Integer := 15;
   ClientsAmount : constant Integer := 10;
   
   MultiplicationMachinesAmount : constant Integer := 8;
   AdditionMachinesAmount : constant Integer := 6;
   MultiplicationDelay : constant Float := 0.7;
   AdditionDelay : constant Float := 0.9;
   PatientEmployerDelay : constant Float := 0.4;

end config;
