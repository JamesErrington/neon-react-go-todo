import { useQuery } from "@tanstack/react-query";
import type { FC } from "react";

function useTodos() {
	return useQuery({
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
				<div key={todo.id}>
					<h4>{todo.task}</h4>
				</div>
			))}
		</div>
	);
};
