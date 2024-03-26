import { useMutation, useQueryClient } from "@tanstack/react-query";
import { type FC, useState } from "react";

export const NewTodoBox: FC = () => {
	const queryClient = useQueryClient();
	const [text, setText] = useState("");
	const mutation = useMutation({
		mutationFn: (value: string) =>
			fetch("http://localhost:8080/todos", {
				method: "POST",
				body: JSON.stringify({ task: value }),
				headers: { "Content-Type": "application/json" },
			}),
		onSuccess: () => {
			queryClient.invalidateQueries({ queryKey: ["todos"] });
			setText("");
		},
	});

	return (
		<div>
			<input
				type="text"
				value={text}
				onChange={(e) => setText(e.target.value)}
			/>
			<button
				type="button"
				disabled={text.length === 0}
				onClick={() => mutation.mutate(text)}
			>
				Create
			</button>
		</div>
	);
};
