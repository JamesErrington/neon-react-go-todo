import { useMutation, useQueryClient } from "@tanstack/react-query";
import type { FC } from "react";

interface Props {
	todo: Todo;
}

export const Todo: FC<Props> = ({ todo }) => {
	const queryClient = useQueryClient();
	const mutation = useMutation({
		mutationFn: (todo: { id: number; task: string; completed: boolean }) =>
			fetch(`http://localhost:8080/todo/${todo.id}`, {
				method: "PUT",
				body: JSON.stringify(todo),
				headers: { "Content-Type": "application/json" },
			}),
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: ["todos"] });
		},
	});

	function handleChange() {
		const newTodo = { ...todo, completed: !todo.completed };
		mutation.mutate(newTodo);
	}

	return (
		<div className={`todo-container ${todo.completed ? "completed" : ""}`}>
			<input type="checkbox" checked={todo.completed} onChange={handleChange} />
			<h4>{todo.task}</h4>
		</div>
	);
};
