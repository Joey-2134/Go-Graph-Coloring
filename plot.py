from __future__ import annotations

import argparse
import csv
import matplotlib.pyplot as plt


def load_run_csv(path):
    iterations: list[int] = []
    conflicts: list[float] = []

    with open(path, "r", newline="", encoding="utf-8") as f:
        reader = csv.DictReader(f)
        required = {"iteration", "conflicts"}
        for row in reader:
            iterations.append(int(row["iteration"]))
            conflicts.append(float(row["conflicts"]))

    return iterations, conflicts


def plot(iterations, conflicts):
    plt.figure(figsize=(8, 4.5))
    plt.plot(iterations, conflicts, marker="o", linewidth=1.5)
    plt.xlabel("Iteration")
    plt.ylabel("Conflicts")
    plt.grid(True, alpha=0.3)
    plt.title("Conflicts per Iteration")

    plt.tight_layout()
    plt.savefig("conflicts.png", bbox_inches="tight")
    plt.close()

if __name__ == "__main__":
    iterations, conflicts = load_run_csv("conflicts.csv")
    plot(iterations, conflicts)

