from __future__ import annotations

import csv
from pathlib import Path

import matplotlib.pyplot as plt


def load_run_csv(path: Path):
    iterations: list[int] = []
    conflicts: list[float] = []

    with open(path, "r", newline="", encoding="utf-8") as f:
        reader = csv.DictReader(f)
        for row in reader:
            iterations.append(int(row["iteration"]))
            conflicts.append(float(row["conflicts"]))

    return iterations, conflicts


def load_all_runs(pattern: str = "conflicts_p_*.csv"):
    runs: list[tuple[str, list[int], list[float]]] = []
    for path in sorted(Path(".").glob(pattern)):
        p_label = path.stem.replace("conflicts_p_", "p=")
        iterations, conflicts = load_run_csv(path)
        runs.append((p_label, iterations, conflicts))
    return runs


def plot(runs):
    if not runs:
        raise FileNotFoundError("No CSV files found matching conflicts_p_*.csv")

    plt.figure(figsize=(8, 4.5))
    for p_label, iterations, conflicts in runs:
        plt.plot(iterations, conflicts, marker="o", linewidth=1.5, label=p_label)

    plt.xlabel("Iteration")
    plt.ylabel("Conflicts")
    plt.grid(True, alpha=0.3)
    plt.title("Conflicts per Iteration by p")
    plt.legend()

    plt.tight_layout()
    plt.savefig("conflicts.png", bbox_inches="tight")
    plt.close()


if __name__ == "__main__":
    runs = load_all_runs()
    plot(runs)

