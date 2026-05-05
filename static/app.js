document.addEventListener('DOMContentLoaded', () => {
    const exerciseList = document.getElementById('exerciseList');
    const exerciseSearch = document.getElementById('exerciseSearch');
    const currentExerciseName = document.getElementById('currentExerciseName');
    const terminal = document.getElementById('terminal');
    const runBtn = document.getElementById('runBtn');
    const clearBtn = document.getElementById('clearBtn');
    
    const taskSection = document.getElementById('taskSection');
    const taskHeader = document.getElementById('taskHeader');
    const taskText = document.getElementById('taskText');
    
    // Toggles
    const verboseToggle = document.getElementById('verboseToggle');
    const formatToggle = document.getElementById('formatToggle');
    const clippyToggle = document.getElementById('clippyToggle');

    let exercises = [];
    let selectedExercise = null;
    let eventSource = null;

    // Load exercises
    async function loadExercises() {
        try {
            const response = await fetch('/api/exercises');
            exercises = await response.json();
            renderExercises(exercises);
        } catch (err) {
            logToTerminal(`Error loading exercises: ${err.message}`, 'error');
        }
    }

    function renderExercises(list) {
        exerciseList.innerHTML = '';
        list.forEach(ex => {
            const li = document.createElement('li');
            li.className = 'exercise-item';
            if (selectedExercise === ex.name) li.classList.add('active');
            li.textContent = ex.name;
            li.onclick = () => selectExercise(ex.name);
            exerciseList.appendChild(li);
        });
    }

    async function selectExercise(name) {
        selectedExercise = name;
        currentExerciseName.textContent = name;
        document.querySelectorAll('.exercise-item').forEach(el => {
            el.classList.toggle('active', el.textContent === name);
        });

        // Load instructions
        try {
            const response = await fetch(`/api/instructions?exercise=${name}`);
            const data = await response.json();
            taskText.textContent = data.instructions;
            // Auto-expand if instructions are found
            if (data.instructions && !data.instructions.includes("No specific instructions")) {
                taskSection.classList.remove('collapsed');
            } else {
                taskSection.classList.add('collapsed');
            }
        } catch (err) {
            taskText.textContent = "Failed to load instructions.";
        }
    }

    taskHeader.onclick = () => {
        taskSection.classList.toggle('collapsed');
    };

    function logToTerminal(text, type = '') {
        const div = document.createElement('div');
        div.className = `line ${type}`;
        
        // Simple ANSI to HTML (focusing on common colors used in the script)
        const html = text
            .replace(/\033\[0;31m/g, '<span style="color: #ef4444">')
            .replace(/\033\[0;32m/g, '<span style="color: #10b981">')
            .replace(/\033\[0;33m/g, '<span style="color: #f59e0b">')
            .replace(/\033\[0;34m/g, '<span style="color: #3b82f6">')
            .replace(/\033\[0;37m/g, '<span style="color: #f1f5f9">')
            .replace(/\033\[0m/g, '</span>');

        div.innerHTML = html;
        terminal.appendChild(div);
        terminal.scrollTop = terminal.scrollHeight;
    }

    exerciseSearch.oninput = (e) => {
        const query = e.target.value.toLowerCase();
        const filtered = exercises.filter(ex => ex.name.toLowerCase().includes(query));
        renderExercises(filtered);
    };

    runBtn.onclick = () => {
        if (!selectedExercise && !confirm('No exercise selected. Run all tests?')) return;

        if (eventSource) {
            eventSource.close();
        }

        logToTerminal(`\n--- Starting Test: ${selectedExercise || 'All'} ---\n`, 'info');
        
        const params = new URLSearchParams({
            exercise: selectedExercise || '',
            verbose: verboseToggle.checked,
            format: formatToggle.checked,
            clippy: clippyToggle.checked
        });

        eventSource = new EventSource(`/api/run?${params.toString()}`);

        eventSource.addEventListener('output', (e) => {
            logToTerminal(e.data);
        });

        eventSource.addEventListener('error', (e) => {
            if (e.data) {
                logToTerminal(`Error: ${e.data}`, 'error');
            } else if (eventSource.readyState === EventSource.CLOSED) {
                logToTerminal('Connection closed.', 'info');
            }
            eventSource.close();
        });

        eventSource.addEventListener('exit', (e) => {
            logToTerminal(`\n--- ${e.data} ---\n`, 'info');
            eventSource.close();
        });
    };

    clearBtn.onclick = () => {
        terminal.innerHTML = '<div class="line welcome">Terminal cleared. Ready...</div>';
    };

    loadExercises();
});
