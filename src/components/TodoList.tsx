import { useQuery } from "@tanstack/react-query";
import type { FC } from "react";
import { Todo } from "./Todo";

function useTodos() {
	return useQuery<Todo[]>({
		queryKey: ["todos"],
		queryFn: () =>
			fetch("http://localhost:8080/todos").then((res) => res.json()),
		initialData: [],
		throwOnError: true,
		retry: false,
	});
}

export const TodoList: FC = () => {
	const { data } = useTodos();

	return (
		<div>
			{data.map((todo) => (
				<Todo key={todo.id} todo={todo} />
			))}
		</div>
	);
};
