export class Client {
  constructor() {
    this.basepath = "/oto/";
  }
}
export class GreeterService {
  constructor(client) {
    this.client = client;
  }
  async greet(greetRequest = null) {
    if (greetRequest == null) {
      greetRequest = new GreetRequest();
    }
    const headers = new Headers();
    headers.set("Accept", "application/json");
    headers.set("Content-Type", "application/json");
    if (this.client.headers) {
      await this.client.headers(headers);
    }
    const response = await fetch(this.client.basepath + "GreeterService.Greet", {
      method: "POST",
      headers,
      body: JSON.stringify(greetRequest)
    });
    if (response.status !== 200) {
      throw new Error(`GreeterService.Greet: ${response.status} ${response.statusText}`);
    }
    return response.json().then((json) => {
      if (json.error) {
        throw new Error(json.error);
      }
      return new GreetResponse(json);
    });
  }
}
export class TodoService {
  constructor(client) {
    this.client = client;
  }
  async create(createTodoRequest = null) {
    if (createTodoRequest == null) {
      createTodoRequest = new CreateTodoRequest();
    }
    const headers = new Headers();
    headers.set("Accept", "application/json");
    headers.set("Content-Type", "application/json");
    if (this.client.headers) {
      await this.client.headers(headers);
    }
    const response = await fetch(this.client.basepath + "TodoService.Create", {
      method: "POST",
      headers,
      body: JSON.stringify(createTodoRequest)
    });
    if (response.status !== 200) {
      throw new Error(`TodoService.Create: ${response.status} ${response.statusText}`);
    }
    return response.json().then((json) => {
      if (json.error) {
        throw new Error(json.error);
      }
      return new CreateTodoResponse(json);
    });
  }
  async delete(deleteTodoRequest = null) {
    if (deleteTodoRequest == null) {
      deleteTodoRequest = new DeleteTodoRequest();
    }
    const headers = new Headers();
    headers.set("Accept", "application/json");
    headers.set("Content-Type", "application/json");
    if (this.client.headers) {
      await this.client.headers(headers);
    }
    const response = await fetch(this.client.basepath + "TodoService.Delete", {
      method: "POST",
      headers,
      body: JSON.stringify(deleteTodoRequest)
    });
    if (response.status !== 200) {
      throw new Error(`TodoService.Delete: ${response.status} ${response.statusText}`);
    }
    return response.json().then((json) => {
      if (json.error) {
        throw new Error(json.error);
      }
      return new DeleteTodoResponse(json);
    });
  }
  async get(getTodoRequest = null) {
    if (getTodoRequest == null) {
      getTodoRequest = new GetTodoRequest();
    }
    const headers = new Headers();
    headers.set("Accept", "application/json");
    headers.set("Content-Type", "application/json");
    if (this.client.headers) {
      await this.client.headers(headers);
    }
    const response = await fetch(this.client.basepath + "TodoService.Get", {
      method: "POST",
      headers,
      body: JSON.stringify(getTodoRequest)
    });
    if (response.status !== 200) {
      throw new Error(`TodoService.Get: ${response.status} ${response.statusText}`);
    }
    return response.json().then((json) => {
      if (json.error) {
        throw new Error(json.error);
      }
      return new GetTodoResponse(json);
    });
  }
  async getAll(getAllTodosRequest = null) {
    if (getAllTodosRequest == null) {
      getAllTodosRequest = new GetAllTodosRequest();
    }
    const headers = new Headers();
    headers.set("Accept", "application/json");
    headers.set("Content-Type", "application/json");
    if (this.client.headers) {
      await this.client.headers(headers);
    }
    const response = await fetch(this.client.basepath + "TodoService.GetAll", {
      method: "POST",
      headers,
      body: JSON.stringify(getAllTodosRequest)
    });
    if (response.status !== 200) {
      throw new Error(`TodoService.GetAll: ${response.status} ${response.statusText}`);
    }
    return response.json().then((json) => {
      if (json.error) {
        throw new Error(json.error);
      }
      return new GetAllTodosResponse(json);
    });
  }
  async update(updateTodoRequest = null) {
    if (updateTodoRequest == null) {
      updateTodoRequest = new UpdateTodoRequest();
    }
    const headers = new Headers();
    headers.set("Accept", "application/json");
    headers.set("Content-Type", "application/json");
    if (this.client.headers) {
      await this.client.headers(headers);
    }
    const response = await fetch(this.client.basepath + "TodoService.Update", {
      method: "POST",
      headers,
      body: JSON.stringify(updateTodoRequest)
    });
    if (response.status !== 200) {
      throw new Error(`TodoService.Update: ${response.status} ${response.statusText}`);
    }
    return response.json().then((json) => {
      if (json.error) {
        throw new Error(json.error);
      }
      return new UpdateTodoResponse(json);
    });
  }
}
export class CreateTodoRequest {
  constructor(data) {
    if (data) {
      this.title = data.title;
      this.description = data.description;
      this.completed = data.completed;
    }
  }
}
export class Todo {
  constructor(data) {
    if (data) {
      this.id = data.id;
      this.title = data.title;
      this.description = data.description;
      this.completed = data.completed;
      this.createdAt = data.createdAt;
      this.updatedAt = data.updatedAt;
    }
  }
}
export class CreateTodoResponse {
  constructor(data) {
    if (data) {
      this.todo = new Todo(data.todo);
      this.error = data.error;
    }
  }
}
export class DeleteTodoRequest {
  constructor(data) {
    if (data) {
      this.todoID = data.todoID;
    }
  }
}
export class DeleteTodoResponse {
  constructor(data) {
    if (data) {
      this.success = data.success;
      this.error = data.error;
    }
  }
}
export class GetAllTodosRequest {
  constructor(data) {
    if (data) {
      this.completed = data.completed;
    }
  }
}
export class GetAllTodosResponse {
  constructor(data) {
    if (data) {
      if (data.todos) {
        this.todos = new Array();
        for (let i = 0; i < data.todos.length; i++) {
          this.todos.push(new Todo(data.todos[i]));
        }
      }
      this.error = data.error;
    }
  }
}
export class GetTodoRequest {
  constructor(data) {
    if (data) {
      this.id = data.id;
    }
  }
}
export class GetTodoResponse {
  constructor(data) {
    if (data) {
      this.todo = new Todo(data.todo);
      this.error = data.error;
    }
  }
}
export class GreetRequest {
  constructor(data) {
    if (data) {
      this.name = data.name;
    }
  }
}
export class GreetResponse {
  constructor(data) {
    if (data) {
      this.greeting = data.greeting;
      this.error = data.error;
    }
  }
}
export class UpdateTodoRequest {
  constructor(data) {
    if (data) {
      this.todo = new Todo(data.todo);
    }
  }
}
export class UpdateTodoResponse {
  constructor(data) {
    if (data) {
      this.todo = new Todo(data.todo);
      this.error = data.error;
    }
  }
}
