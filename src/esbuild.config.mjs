import { build } from "esbuild"
import { copy } from "esbuild-plugin-copy"

const BUILD_DIR = "build"
const isProduction = process.env.NODE_ENV === "production";

build({
	entryPoints: ["src/index.tsx"],
	bundle: true,
	format: "esm",
	minify: isProduction,
	sourcemap: !isProduction,
	outfile: `${BUILD_DIR}/index.js`,
	logLevel: "info",
	color: true,
	plugins: [
		copy({
			resolveFrom: 'cwd',
			assets: {
				from: ['./public/*'],
				to: [`./${BUILD_DIR}`],
			},
		}),
	],
});
