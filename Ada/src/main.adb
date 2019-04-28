with Ada.Text_IO;    use Ada.Text_IO;
with Ada.Integer_Text_IO; use Ada.Integer_Text_IO;
with Ada.Strings.Unbounded; use Ada.Strings.Unbounded;
with Ada.Numerics.Discrete_Random;
with config;

procedure Main is

   type Item is record
      Value: Integer;
   end record;

   type TaskToDo is record
      FirstArgument: Integer;
      SecondArgument: Integer;
      Operation: Character;
      Result: Integer;
   end record;

   type Statistics is
      record
         id : Integer;
         EmployerType : Unbounded_String;
         counter : Integer;
      end record;

   type MachineStatistics is
      record
         id : Integer;
         counter : Integer;
      end record;

   type ListofTasks is array (0.. config.TasksListSize) of TaskToDo;
   type ListofItems is array (0.. config.ItemsSize) of Item;
   type EmployerType is array (0 .. 1) of Unbounded_String;
   EmployersStatistics : array (0 .. config.EmployersAmount-1) of Statistics;
   AdditionMachineStatistics : array (0 .. config.AdditionMachinesAmount-1) of MachineStatistics;
   MultiplicationMachineStatistics : array (0 .. config.MultiplicationMachinesAmount-1) of MachineStatistics;
   EmployerTypes : EmployerType := (To_Unbounded_String("patient"), To_Unbounded_String("impatient"));

   task ToDoList is
      entry Insert (An_Item : in TaskToDo);
      entry Remove (An_Item : out TaskToDo);
      entry Show (A_List : out ListofTasks);
   end ToDoList;

   task Magazine is
      entry Insert (An_Item : in Item);
      entry Remove (An_Item : out Item);
      entry Show (A_List : out ListofItems);
   end Magazine;

   task type CEO;
   task type Employer is
      entry start(Id : in Integer);
   end Employer;

   task type Client is
      entry start(Id : in Integer);
   end Client;

   task type AdditionMachine is
      entry start(Id : in Integer);
      entry Insert (An_Item : in TaskToDo; Out_Item : out TaskToDo);
   end AdditionMachine;

   task type MultiplicationMachine is
      entry start(Id : in Integer);
      entry Insert (An_Item : in TaskToDo; Out_Item : out TaskToDo);
   end MultiplicationMachine;

   function addition (A, B : Integer) return Integer is
   begin
      return A+B;
   end addition;

   function multiplication (A, B : Integer) return Integer is
   begin
      return A*B;
   end multiplication;

   function computeTask(t : in TaskToDo) return Integer is
   operation : Character;
   value : Integer;
   begin
      operation := t.Operation;
      if operation = 'a' then
         value := addition(t.FirstArgument, t.SecondArgument);
      else
         value := multiplication(t.FirstArgument, t.SecondArgument);
      end if;
      return value;
   end computeTask;

   function createNewTask return TaskToDo is
      subtype Operation_Range is Integer range 0 .. 1;
      subtype Random_Range is Integer range 1 .. 1000;
      package RandomOperation is new Ada.Numerics.Discrete_Random(Operation_Range);
      package RandomTask is new Ada.Numerics.Discrete_Random(Random_Range);

      type OperationArray is array (0 .. 1) of Character;
      TaskOperations : OperationArray := ('a' , 'm');

      newTask : TaskToDo;
      seedOptions : RandomOperation.Generator;
      seedArguments : RandomTask.Generator;
      firstArgument : Integer;
      secondArgument : Integer;
      operation : Character;

   begin
      RandomOperation.Reset (seedOptions);
      RandomTask.Reset (seedArguments);
      firstArgument := RandomTask.Random (seedArguments);
      secondArgument := RandomTask.Random (seedArguments);
      operation := TaskOperations(RandomOperation.Random (seedOptions));
      newTask := (firstArgument, secondArgument, operation, 0);

      return newTask;

   end createNewTask;

   simulationMode: Character := 'c';

   task body ToDoList is
      limit : Integer := config.TasksListSize;
      length : Integer:= 0;
      MyList : ListofTasks;
      Head: Integer := 0;
      Tail: Integer := 0;

   begin
      loop
      select
         when length < limit =>
            accept Insert (An_Item : in TaskToDo) do
               MyList(Tail) := An_Item;
            end Insert;
            Tail := (Tail+1) mod limit;
            length := length+1;
      or
         when length > 0 => accept Remove (An_Item : out TaskToDo) do
              An_Item := MyList(Head);
            end Remove;
            Head := (Head + 1) mod limit;
            length := length-1;
      or
         accept Show (A_List : out ListofTasks) do
               A_list := MyList;
               Put_Line("Below you can see the current state of ToDoList");
              for I in 0 .. length-1 loop
                  Put_Line(MyList(I).FirstArgument'Image & " " & MyList(I).SecondArgument'Image & " " & MyList(I).Operation'Image);
               end loop;
               Put_Line("");
         end Show;
         end select;
      end loop;

   end ToDoList;

   task body Magazine is
      limit : Integer := config.ItemsSize;
      length : Integer:= 0;
      MyList : ListofItems;
      Head: Integer := 0;
      Tail: Integer := 0;

   begin
      loop
      select
         when length < limit =>
            accept Insert (An_Item : in Item) do
               MyList(Tail) := An_Item;
            end Insert;
            Tail := (Tail+1) mod limit;
            length := length+1;
      or
         when length > 0 => accept Remove (An_Item : out Item) do
              An_Item := MyList(Head);
            end Remove;
            Head := (Head + 1) mod limit;
            length := length-1;
      or
         accept Show (A_List : out ListofItems) do
               A_list := MyList;
               Put_Line("Below you can see the current state of Magazine");
               for I in MyList'Range loop
                  Put_Line(MyList(I).Value'Image);
               end loop;
               Put_Line("");

         end Show;
         end select;

      end loop;



end Magazine;

   procedure taskToString(TaskToPrint : in TaskToDo) is
   begin
      Put(Integer'Image(TaskToPrint.FirstArgument) & " " & Integer'Image(TaskToPrint.FirstArgument) & " " & (TaskToPrint.Operation) & " " & Integer'Image(TaskToPrint.Result));
      New_Line;
   end taskToString;

task body MultiplicationMachine is
      multiTask : TaskToDo;
      MachineId : Integer;
      Busy : Boolean := False;
      Result : Integer;
      Counter: Integer;
      CurrentMultiplicationMachineStatisctic : MachineStatistics;
   begin
       accept start(Id : in Integer) do
         MachineId := Id;
         Counter:=0;
         CurrentMultiplicationMachineStatisctic := (MachineId, Counter);
         MultiplicationMachineStatistics(MachineId-1) := CurrentMultiplicationMachineStatisctic;
       end start;
       loop
         select
            when Busy = False =>
               accept Insert (An_Item : in TaskToDo; Out_Item : out TaskToDo) do
                  Busy := True;

                  Result := multiplication(An_Item.FirstArgument, An_Item.SecondArgument);
                  delay Duration(config.MultiplicationDelay);
                  multiTask := (An_Item.FirstArgument, An_Item.SecondArgument, An_Item.Operation, Result);
                  Out_Item := multiTask;
                  Busy := False;
                  if simulationMode = 'v' then
                     Put("Multiplication Machine " & Integer'Image(MachineId) & " finished the task");
                     taskToString(multiTask);
                  end if;
                  Counter := Counter + 1;
                  MultiplicationMachineStatistics(MachineId-1).counter := Counter;
               end Insert;
         end select;
      end loop;
end MultiplicationMachine;


 task body AdditionMachine is
      addTask : TaskToDo;
      MachineId : Integer;
      Busy : Boolean := False;
      Result : Integer;
      Counter: Integer;
      CurrentAdditionMachineStatisctic : MachineStatistics;
   begin
       accept start(Id : in Integer) do
         MachineId := Id;
         Counter := 0;
         CurrentAdditionMachineStatisctic := (MachineId, Counter);
         AdditionMachineStatistics(MachineId-1) := CurrentAdditionMachineStatisctic;
       end start;
       loop
         select
            when Busy = False =>
               accept Insert (An_Item : in TaskToDo; Out_Item : out TaskToDo) do
                  Busy := True;
                  Result := addition(An_Item.FirstArgument, An_Item.SecondArgument);
                  delay Duration(config.AdditionDelay);
                  addTask := (An_Item.FirstArgument, An_Item.SecondArgument, An_Item.Operation, Result);
                  Out_Item := addTask;
                  Busy := False;
                  if simulationMode = 'v' then
                     Put("Addition Machine " & Integer'Image(MachineId) & " finished the task");
                     taskToString(addTask);
                  end if;
                  Counter := Counter + 1;
                  AdditionMachineStatistics(MachineId-1).counter := Counter;
               end Insert;
         end select;
      end loop;
   end AdditionMachine;

   function whoAmI return Unbounded_String is
      subtype Random_Range is Integer range 0 .. 1;
      package R is new Ada.Numerics.Discrete_Random (Random_Range);

      G : R.Generator;
      EmployerType : Unbounded_String;
     begin
      R.Reset (G);
      EmployerType := EmployerTypes(R.Random (G));

     return EmployerType;
   end whoAmI;


   function chooseMultiplicationMachine return Integer is
      subtype Random_Range is Integer range 1 .. config.MultiplicationMachinesAmount;
      package R is new Ada.Numerics.Discrete_Random (Random_Range);
      G : R.Generator;
     begin
      R.Reset (G);
     return R.Random (G);
   end chooseMultiplicationMachine;


   function chooseAdditionMachine return Integer is
      subtype Random_Range is Integer range 1 .. config.AdditionMachinesAmount;
      package R is new Ada.Numerics.Discrete_Random (Random_Range);
      G : R.Generator;
     begin
      R.Reset (G);
     return R.Random (G);
   end chooseAdditionMachine;


   type AdditionMachines is array (Positive range 1 .. config.AdditionMachinesAmount) of AdditionMachine;
   AdditionMachinesList : AdditionMachines;

   type MultiplicationMachines is array (Positive range 1 .. config.MultiplicationMachinesAmount) of MultiplicationMachine;
   MultiplicationMachinesList : MultiplicationMachines;



   task body Client is
   takenItem : Item;
   clientId : Integer;
begin
   accept start(Id : in Integer) do
         clientId := Id;
      end start;
      Endless_Loop :
      loop
          delay Duration(config.ClientDelay);
         Magazine.Remove(takenItem);
         if simulationMode = 'v' then
            Put_Line("Client " & Integer'Image(clientId) & " bought the new product" & Integer'Image(takenItem.Value));
         end if;


      end loop Endless_Loop;
   end Client;

   task body CEO is
      newTask : TaskToDo;
   begin
      Endless_Loop :
      loop
         delay Duration(config.CeoDelay);
         newTask := createNewTask;

         ToDoList.Insert(newTask);

         if simulationMode = 'v' then
            Put_Line("CEO add new Task: " & Integer'Image(newTask.FirstArgument) & " " & Integer'Image(newTask.SecondArgument) & " " & (newTask.Operation));
         end if;

      end loop Endless_Loop;


   end CEO;

task body Employer is
   takenTask : TaskToDo;
   resultTask : TaskToDo;
   result : Integer;
   EmployerId: Integer;
   MachineId : Integer;
   Done : Boolean;
   newItem : Item;
   TypeofEmployer : Unbounded_String;
   Counter : Integer;
   CurrentEmployerStatisctic : Statistics;
   begin
       accept start(Id : in Integer) do
         EmployerId := Id;
         TypeofEmployer := whoAmI;
         Counter := 0;
         CurrentEmployerStatisctic := (EmployerId, TypeofEmployer, Counter);
         EmployersStatistics(EmployerId-1) := CurrentEmployerStatisctic;
      end start;
      Endless_Loop :
      loop
         ToDoList.remove(takenTask);
         if simulationMode = 'v' then
            Put_Line("Employer " & Integer'Image(EmployerId) & " got task " &  Integer'Image(takenTask.FirstArgument) & " " & Integer'Image(takenTask.SecondArgument) & " " & (takenTask.Operation) );
         end if;
          if takenTask.Operation = 'a' then
            MachineId := chooseAdditionMachine;
         else
            MachineId := chooseMultiplicationMachine;
         end if;

         if TypeofEmployer = "impatient" then
            Done := false;
            if takenTask.Operation = 'a' then
              while not Done loop
                 select
                     AdditionMachinesList(MachineId).Insert(takenTask, resultTask);
                     Done := True;
                  or
                       delay Duration(config.PatientEmployerDelay);
                       MachineId := chooseAdditionMachine;
                  end select;
               end loop;
            else
               while not Done loop
                 select
                     MultiplicationMachinesList(MachineId).Insert(takenTask, resultTask);
                     Done := True;
                  or
                       delay Duration(config.PatientEmployerDelay);
                       MachineId := chooseMultiplicationMachine;
                  end select;
               end loop;
            end if;
         else
            if takenTask.Operation = 'a' then
               AdditionMachinesList(MachineId).Insert(takenTask, resultTask);
            else
               AdditionMachinesList(MachineId).Insert(takenTask, resultTask);
            end if;
         end if;

         result := resultTask.Result;
         if simulationMode = 'v' and resultTask.Operation = 'm' then
            Put("Employer " & Integer'Image(EmployerId) & " took from Multiplication Machine " & Integer'Image(MachineId) & " task: ");
            taskToString(resultTask);
         end if;
         if simulationMode = 'v' and resultTask.Operation = 'a' then
            Put("Employer " & Integer'Image(EmployerId) & " took from Addition Machine "& Integer'Image(MachineId) & " task:" );
            taskToString(resultTask);
         end if;
         Counter := Counter + 1;
         EmployersStatistics(EmployerId-1).counter := Counter;
         newItem.Value := result;
         Magazine.Insert(newItem);
         delay Duration(config.EmployerDelay);
      end loop Endless_Loop;

   end Employer;


   optionCalmMode: Character;

   procedure calmMode is
      magazineState: ListofItems;
      toDoListState: ListofTasks;
   begin
      Endless_Loop:
      loop
         Put_Line("");
         Put_Line("");
         Put_Line("The simulator is running in the calm mode.");
	 Put_Line( "You can get the information about the current state of simulation by typing:");
	 Put_Line("1 - information about the current magazine state");
         Put_Line("2 - information about the current task to do");
         Put_Line("3 - information about the employees");
         Put_Line("4 - information about the machines");
         Get(optionCalmMode);
         if optionCalmMode = '1' then
            Magazine.Show(magazineState);
         elsif optionCalmMode = '2' then
            ToDoList.Show(toDoListState);
         elsif optionCalmMode = '3' then
            Put_Line("Below you can see the current statistics of employees");
            for i in EmployersStatistics'Range loop
               Put_Line("Employee Id :" & Integer'Image(EmployersStatistics(i).id) & "  type : " &
                          To_String(EmployersStatistics(i).EmployerType) & "  successes :" &
                          Integer'Image(EmployersStatistics(i).counter));
            end loop;
         elsif optionCalmMode = '4' then
            Put_Line("Below you can see the current statistics of addition Machines");
            for i in AdditionMachineStatistics'Range loop
               Put_Line("Machine Id :" & Integer'Image(AdditionMachineStatistics(i).id) & " amount of resolved Tasks :" &
                          Integer'Image(AdditionMachineStatistics(i).counter));
            end loop;
            Put_Line("");
            Put_Line("Below you can see the current statistics of multiplication Machines");
            for i in MultiplicationMachineStatistics'Range loop
               Put_Line("Machine Id :" & Integer'Image(MultiplicationMachineStatistics(i).id) & " amount of resolved Tasks :" &
                          Integer'Image(MultiplicationMachineStatistics(i).counter));
            end loop;
         end if;



      end loop Endless_Loop;



   end calmMode;



   c: CEO;

   type Employers is array (Positive range 1 .. config.EmployersAmount) of Employer;
   EmployerList : Employers;

   type Clients is array (Positive range 1 .. config.ClientsAmount) of Client;
   ClientList : Clients;

begin

   for i in MultiplicationMachinesList'Range loop
      MultiplicationMachinesList(i).start(i);
   end loop;

   for i in AdditionMachinesList'Range loop
      AdditionMachinesList(i).start(i);
   end loop;

   for i in EmployerList'Range loop
      EmployerList(i).start(i);
   end loop;

   for i in ClientList'Range loop
      ClientList(i).start(i);
   end loop;


   Endless_Loop:
   loop
      Put_Line("Choose a simulator mode you'd like to work with ");
      Put_Line("c - calm mode ");
      Put_Line("v - verbose mode ");
      Get(simulationMode);
      if simulationMode = 'v' then
         exit Endless_Loop;
      elsif simulationMode = 'c' then
         calmMode;
      end if;
   end loop Endless_Loop;
end Main;
