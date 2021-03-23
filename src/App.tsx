import { h } from "preact";
import { useEffect, useState, useCallback } from "preact/hooks";
import { Create } from "./Create";
import { Client, Todo, TodoService } from "./types/oto.gen";
export const client = new Client();

function App() {
  const [todos, setTodos] = useState<Todo[]>([]);

  const fetchData = useCallback(async () => {
    const service = new TodoService(client);

    const tt = await service.getAll();
    setTodos(tt.todos ?? []);
  }, []);

  useEffect(() => {
    fetchData();
  }, []);

  const onSubmit = () => {
    fetchData();
  };
  return (
    <div className="App">
      <Create onSubmit={onSubmit} />
      {!todos ? (
        <p>Could not fetch todos</p>
      ) : !todos.length ? (
        <p>No todos created</p>
      ) : (
        todos.map((t) => <div key={t.id}>{t.title}</div>)
      )}
    </div>
  );
}

export default App;
