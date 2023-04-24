USE empresa_db;

/* ________________________________________________________________________________________________________________________ */
/* Seleccionar el nombre, el puesto y la localidad de los departamentos donde trabajan los vendedores. */
SELECT e.nombre, e.puesto, d.localidad FROM employee as e JOIN departamento as d ON e.depto_nro = d.depto_nro
WHERE e.puesto = "Vendedor";


/* ________________________________________________________________________________________________________________________ */
/* Visualizar los departamentos con más de cinco empleados. */
SELECT d.* FROM departamento as d WHERE (SELECT COUNT(*) FROM employee as e WHERE e.depto_nro = d.depto_nro) >= 2;

SELECT d.*, COUNT(*) as cantidad FROM employee as e INNER JOIN departamento as d ON e.depto_nro = d.depto_nro
GROUP BY e.depto_nro HAVING cantidad > 1;


/* ________________________________________________________________________________________________________________________ */
/* Mostrar el nombre, salario y nombre del departamento de los empleados que tengan el mismo puesto que ‘Mito Barchuk’ */
SELECT e.nombre, e.salario, d.nombre_depto FROM employee as e JOIN departamento as d ON e.depto_nro = d.depto_nro
WHERE e.puesto = (
	SELECT e2.puesto FROM employee as e2 WHERE e2.nombre = 'Mito' AND e2.apellido = 'Barchuk'
);


/* ________________________________________________________________________________________________________________________ */
/* Mostrar los datos de los empleados que trabajan en el departamento de contabilidad, ordenados por nombre. */
SELECT e.* FROM employee as e
WHERE e.depto_nro = (
	SELECT d.depto_nro FROM departamento as d WHERE d.nombre_depto = 'Contabilidad'
)
ORDER BY e.nombre DESC;


/* ________________________________________________________________________________________________________________________ */
/* Mostrar el nombre del empleado que tiene el salario más bajo. */
SELECT e.nombre FROM employee as e
WHERE e.salario = (
	SELECT MIN(e2.salario) FROM employee as e2
);


/* ________________________________________________________________________________________________________________________ */
/* Mostrar los datos del empleado que tiene el salario más alto en el departamento de ‘Ventas’. */
SELECT e.* FROM employee as e
WHERE e.depto_nro = (
	SELECT d.depto_nro FROM departamento as d WHERE d.nombre_depto = 'Ventas'
)
AND e.salario = (
	SELECT MAX(e2.salario) FROM employee as e2
	WHERE e2.depto_nro = (
		SELECT d.depto_nro FROM departamento as d WHERE d.nombre_depto = 'Ventas'
	)
);

SELECT e.* FROM employee as e JOIN departamento as d ON e.depto_nro = d.depto_nro
WHERE d.nombre_detpo = "Ventas"
AND e.salario = (
	SELECT MAX(e2.salario) FROM employee as e2
    WHERE e2.depto_nro = d.depto_nro
)