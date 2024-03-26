import { QueryClient, QueryClientProvider } from "@tanstack/react-query";
import type { FC } from "react";
import { NewTodoBox } from "./components/NewTodoBox";
import { TodoList } from "./components/TodoList";

const queryClient = new QueryClient();

export const App: FC = () => {
	return (
		<QueryClientProvider client={queryClient}>
			<NewTodoBox />
			<TodoList />
		</QueryClientProvider>
	);
};
