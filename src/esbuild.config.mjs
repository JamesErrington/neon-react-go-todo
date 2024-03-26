import esbuild from "esbuild";
import { copy } from "esbuild-plugin-copy";

const BUILD_DIR = "build";
const isProduction = process.env.NODE_ENV === "production";
const watch = process.argv.includes("--watch");

const config = {
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
			resolveFrom: "cwd",
			assets: {
				from: ["./public/*"],
				to: [`./${BUILD_DIR}`],
			},
		}),
	],
};

if (watch) {
	const ctx = await esbuild.context(config);
	await ctx.watch();
} else {
	esbuild.build(config);
}
