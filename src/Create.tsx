import { h } from "preact";
import { useState } from "preact/hooks";
import { client } from "./App";
import { TodoService } from "./types/oto.gen";

interface Props {
  onSubmit: VoidFunction;
}

function Create({ onSubmit }: Props) {
  const [title, setTitle] = useState("");

  const handleSubmit = async () => {
    const service = new TodoService(client);
    const res = await service.create({
      title: title,
      description: "",
      completed: false,
    });
    if (res.error) alert(res.error);
    onSubmit();
  };
  return (
    <div>
      <input
        type="text"
        name="title"
        onChange={(e) => setTitle(e.currentTarget.value)}
      ></input>

      <button type="button" onClick={handleSubmit}>
        Submit
      </button>
    </div>
  );
}

export { Create };
